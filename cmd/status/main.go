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

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	defer tw.Flush()

	fmt.Fprint(tw, "account\tlast tx\tdays\tlast file\n")
	fmt.Fprint(tw, "-------\t-------\t----\t---------\n")

	var accounts []string
	for name, _ := range config.Account {
		accounts = append(accounts, name)
	}

	sort.Strings(accounts)

	for _, name := range accounts {
		acc := config.Account[name]

		// parse dirs
		matches, err := util.MatchDir(acc, config.Pfin.Root)
		if err != nil {
			log.Fatal(err)
		}

		// parse transactions
		tx, err := util.ParseDir(acc, config.Pfin.Root)
		if err != nil {
			log.Fatal(err)
		}

		sort.SliceStable(tx, func(i, j int) bool {
			return tx[i].Date().Before(tx[j].Date())
		})

		// output
		last := tx[len(tx)-1].Date()
		fmt.Fprintf(
			tw, "%s\t%s\t%d\t%s\n",
			name,
			util.FormatDate(last),
			int(time.Now().Sub(last).Hours())/24,
			filepath.Base(matches[len(matches)-1]),
		)
	}
}
