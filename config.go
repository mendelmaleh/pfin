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
	// parser type
	Type string

	// map of users to card identifier
	Users map[string]string
}

func (a Account) User(card string) string {
	for k, v := range a.Users {
		if v == card {
			return k
		}
	}

	return card
}

func ParseConfig(path string) (config Config, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	if err = toml.Unmarshal(data, &config); err != nil {
		return
	}

	// make root filepath absolute
	config.Pfin.Root = filepath.Clean(filepath.Join(filepath.Dir(path), config.Pfin.Root))

	// make a slice of account keys
	// TODO: figure out where/why (just determinism?)
	/*
		for k, v := range config.Account {
			// set type to name if not set
			if v.Type == "" {
				v.Type = k
				config.Account[k] = v
			}

			// config.Accounts = append(config.Accounts, k)
		}
	*/
	// sort.Strings(config.Accounts)

	return
}
