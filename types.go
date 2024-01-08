package pfin

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Pfin struct {
		User string
		Root string
	}

	// the map is nice in the toml config, but not that nice for usability
	Account  map[string]Account
	Accounts []string
}

type Account struct {
	Name string

	// parser type, inherited from account name if unset
	Type string

	// folder path from root, inherited from account name if unset
	Path string

	// default user, inherited from config.Pfin.User if unset
	DefaultUser string `toml:"user"`

	// map of users to card identifier
	Users map[string][]string

	// generated
	Cards map[string]string
}

func (a Account) User(card string) string {
	if user, ok := a.Cards[card]; ok {
		return user
	}

	return a.DefaultUser
}

// ParseConfig will use default config location if path is empty
func ParseConfig(path string) (config Config, err error) {
	// default config path
	if path == "" {
		configpath, err := os.UserConfigDir()
		if err != nil {
			return config, err
		}

		path = filepath.Join(configpath, "pfin", "statements", "config.toml")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	if err = toml.Unmarshal(data, &config); err != nil {
		return
	}

	// make root filepath absolute
	config.Pfin.Root = filepath.Clean(filepath.Join(filepath.Dir(path), config.Pfin.Root))

	for k, v := range config.Account {
		// store the name in the struct and in the list of (sorted) accounts
		v.Name = k
		config.Accounts = append(config.Accounts, k)

		// set type to name if unset
		if v.Type == "" {
			v.Type = k
		}

		// set path to name if unset
		if v.Path == "" {
			v.Path = k
		}

		// set default user if unset
		if v.DefaultUser == "" {
			v.DefaultUser = config.Pfin.User
		}

		// set card to user map
		v.Cards = make(map[string]string)
		for user, cards := range v.Users {
			for _, card := range cards {
				v.Cards[card] = user
			}
		}

		// reassign modified copy to map
		config.Account[k] = v
	}

	sort.Strings(config.Accounts)

	return
}
