package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/capitalone"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/personal"
)

type Opts struct {
	Users, Accounts, Payments util.StringFilter
}

func main() {
	var opts Opts
	flag.StringVar(&opts.Users.String, "user", "", "filter user")
	flag.StringVar(&opts.Accounts.String, "account", "", "filter account")
	flag.StringVar(&opts.Payments.String, "payment", "payments", "filter payments account")

	flag.Parse()

	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	var txns []pfin.Transaction
	var credit float64

	for name, acc := range config.Account {
		if opts.Accounts.Filter(name) && opts.Payments.Filter(name) {
			continue
		}

		dirtxns, err := util.ParseDir(acc, config.Pfin.Root)
		if err != nil {
			log.Fatal(err)
		}

		for _, tx := range dirtxns {
			if opts.Users.Filter(tx.User()) {
				continue
			}

			txns = append(txns, tx)

			if tx.Amount() < 0 {
				credit += tx.Amount()
			}
		}
	}

	credit *= -1
	// fmt.Printf("credit: %.2f\n", credit)

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	for _, tx := range txns {
		a := tx.Amount()

		// print credits and uncovered debits
		if a < 0 || a > credit {
			fmt.Fprintln(tw, pfin.TxString(tx, "\t"))
		}

		// sub debits
		if a > 0 {
			credit -= a
		}
	}

	tw.Flush()
	fmt.Printf("\ntotal: %.2f\n", credit)
}
