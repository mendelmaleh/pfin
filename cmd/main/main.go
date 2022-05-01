package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/parser/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/amex"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/bofa"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/capitalone"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/personal"
)

type Opts struct {
	User    string
	Account string
}

func main() {
	// flags
	var opts Opts
	flag.StringVar(&opts.User, "user", "", "filter user")
	flag.StringVar(&opts.Account, "account", "", "filter account")

	flag.Parse()

	// config
	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal("error parsing config: ", err)
	}

	// collect
	var txns []pfin.Transaction
	for name, acc := range config.Account {
		tx, err := util.ParseDir(acc, filepath.Join(config.Pfin.Root, name))
		if err != nil {
			panic(err)
		}

		txns = append(txns, tx...)
	}

	// sort oldest to newest
	sort.SliceStable(txns, func(i, j int) bool {
		return txns[i].Date().Before(txns[j].Date())
	})

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	var sum = map[string]float64{}

	for _, tx := range txns {
		if opts.User != "" && tx.User() != opts.User {
			continue
		}

		if opts.Account != "" && tx.Account() != opts.Account {
			continue
		}

		if _, ok := sum[tx.User()]; !ok {
			sum[tx.User()] = 0
		}

		sum[tx.User()] += tx.Amount()
		fmt.Fprintln(tw, pfin.TxString(tx, "\t"))
	}

	tw.Flush()
	fmt.Println()

	tw.Init(os.Stdout, 0, 8, 1, '\t', 0)
	for user, total := range sum {
		fmt.Fprintf(tw, "%s\t%.2f\n", user, total)
	}

	tw.Flush()
}
