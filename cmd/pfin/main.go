package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sort"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/all"
)

func main() {
	// flags
	opts := struct {
		parser *string
	}{
		flag.String("parser", "", "parser to use for files"),
	}

	flag.Parse()

	// config
	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	var txns []pfin.Transaction

	// by default parse all files
	if len(flag.Args()) == 0 {
		for _, k := range config.Accounts {
			acc := config.Account[k]

			// using util is a little unnatural, maybe it should be a method of account
			// same with config.Pfin.Root
			tx, err := util.ParseDir(acc, config.Pfin.Root)
			if err != nil {
				log.Fatal(err)
			}

			txns = append(txns, tx...)
		}
	} else {
		acc, ok := config.Account[*opts.parser]
		if !ok {
			log.Fatalf("invalid account %q, should be one of %v", *opts.parser, config.Accounts)
		}

		// deduplicate with parsedir?
		for _, f := range flag.Args() {
			file, err := os.ReadFile(f)
			if err != nil {
				log.Fatal(err)
			}

			// simplify logic, not very intuitive, and should be able to parse without a config and account, too much setup
			tx, err := pfin.Parse(acc, filepath.Base(f), file)
			if err != nil {
				log.Fatal(err)
			}

			// TODO: use copy(), consider parallelizing
			txns = append(txns, tx...)
		}

	}

	// sort oldest to newest
	sort.SliceStable(txns, func(i, j int) bool {
		return txns[i].Date().Before(txns[j].Date())
	})

	// csv
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"date", "amount", "account", "card", "user", "name", "category"})

	// what a fucking mess
	for _, tx := range txns {
		w.Write([]string{
			util.FormatDate(tx.Date()),
			util.FormatCents(tx.Amount()),
			tx.Account(),
			tx.Card(),
			tx.User(),
			tx.Name(),
			tx.Category(),
		})
	}

	w.Flush()
}
