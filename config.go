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

	// set type to name if not set
	for k, v := range config.Account {
		if v.Type == "" {
			v.Type = k
		}

		config.Account[k] = v
	}

	return
}
