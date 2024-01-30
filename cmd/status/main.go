package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/all"
)

func main() {
	// parse config
	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	defer tw.Flush()

	fmt.Fprint(tw, "bank\taccount\tbalance\tdebits\tcredits\tlast tx\tdays\tlast file\n")
	fmt.Fprint(tw, "----\t-------\t-------\t------\t-------\t-------\t----\t---------\n")

	var accounts []string
	for name := range config.Account {
		accounts = append(accounts, name)
	}

	sort.Strings(accounts)
	sort.SliceStable(accounts, func(i, j int) bool {
		return config.Account[accounts[i]].Type < config.Account[accounts[j]].Type
	})

	for _, name := range accounts {
		acc := config.Account[name]

		// parse dirs
		matches, err := util.MatchDir(acc, config.Pfin.Root)
		if err != nil {
			log.Fatal(err)
		}

		// parse transactions
		txns, err := util.ParseDir(acc, config.Pfin.Root)
		if err != nil {
			log.Fatal(err)
		}

		// account totals
		var credits, debits float64
		for _, tx := range txns {
			if a := tx.Amount(); a < 0 {
				credits -= a
			} else {
				debits -= a
			}
		}

		balance := credits + debits

		// find last tx
		sort.SliceStable(txns, func(i, j int) bool {
			return txns[i].Date().Before(txns[j].Date())
		})
		last := txns[len(txns)-1].Date()

		// output
		fmt.Fprintf(
			tw, "%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n",
			acc.Type,
			name,
			util.FormatCents(balance),
			util.FormatCents(debits),
			util.FormatCents(credits),
			util.FormatDate(last),
			int(time.Now().Sub(last).Hours())/24,
			// assumes files are named alphabetically
			filepath.Base(matches[len(matches)-1]),
		)
	}
}
