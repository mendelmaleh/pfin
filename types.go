package pfin

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Pfin struct {
		User string
		Root string
	}

	Account map[string]Account
}

type Account struct {
	// parser type, inherited from account name if unset
	Type string

	// default user, inherited from config.Pfin.User if unset
	DefaultUser string

	// map of users to card identifier
	Users map[string]string
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

		path = filepath.Join(configpath, "pfin", "config.toml")
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
		// set type to name if unset
		if v.Type == "" {
			v.Type = k
		}

		// set default user if unset
		if v.DefaultUser == "" {
			v.DefaultUser = config.Pfin.User
		}

		// set card to user map
		v.Cards = make(map[string]string)
		for user, card := range v.Users {
			v.Cards[card] = user
		}

		config.Account[k] = v
	}

	return
}
