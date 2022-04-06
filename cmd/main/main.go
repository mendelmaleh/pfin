package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/parser/amex"
	"git.sr.ht/~mendelmaleh/pfin/parser/capitalone"
	"git.sr.ht/~mendelmaleh/pfin/parser/util"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Pfin struct {
		User string
		Root string
	}

	Account map[string]Account

	// Account keys/names, sorted, for iteration
	Accounts []string
}

type Account struct {
	Path  string
	Users map[string]string
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath := filepath.Clean(filepath.Join(cwd, "../../config.toml"))

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	// make root filepath absolute
	config.Pfin.Root = filepath.Clean(filepath.Join(filepath.Dir(configPath), config.Pfin.Root))

	// make a slice of account keys
	for k, _ := range config.Account {
		config.Accounts = append(config.Accounts, k)
	}
	sort.Strings(config.Accounts)

	// parsing functions, key to fn
	parsefns := map[string]func([]byte, *[]pfin.Transaction) error{
		"amex": amex.Parse,
		// "bofa":       bofa.Parse,
		"capitalone": capitalone.Parse,
	}

	var txns []pfin.Transaction
	for _, name := range config.Accounts {
		fn, ok := parsefns[name]
		if !ok {
			fmt.Printf("skipping %q, no parsing function defined\n", name)
			continue
		}

		tx, err := util.ParseDir(filepath.Join(config.Pfin.Root, name), fn)
		if err != nil {
			panic(err)
		}

		txns = append(txns, tx...)
	}

	// sort oldest to newest
	sort.SliceStable(txns, func(i, j int) bool {
		return txns[i].Date().Before(txns[j].Date())
	})

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	defer tw.Flush()

	for _, tx := range txns {
		io.WriteString(tw, pfin.TxString(tx, "\t"))
		tw.Write([]byte{'\n'})
		// fmt.Fprintf(tw, "%s\t%.2f\t%s\n", tx.Date().Format(pfin.ISO8601), tx.Amount(), tx.Name())
		// fmt.Printf("%+v\n", v)
		// spew.Dump(v)
	}
}
