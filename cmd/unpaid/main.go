package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/capitalone"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/personal"
)

type Opts struct {
	Users, Accounts, Payments util.StringFilter

	Separator string
}

func main() {
	var opts Opts

	flag.StringVar(&opts.Users.String, "user", "", "filter user")
	flag.StringVar(&opts.Accounts.String, "account", "", "filter account")
	flag.StringVar(&opts.Payments.String, "payment", "payments", "filter payments account")

	flag.StringVar(&opts.Separator, "sep", "\t", "separator")

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

	sort.SliceStable(txns, func(i, j int) bool {
		return txns[i].Date().Before(txns[j].Date())
	})

	credit *= -1
	// fmt.Printf("credit: %.2f\n", credit)

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	var total, paid float64
	for _, tx := range txns {
		a := tx.Amount()

		// print credits and uncovered debits
		if a < 0 || a > credit {
			total += a

			fmt.Fprintln(tw, strings.Join([]string{
				tx.Date().Format(pfin.ISO8601),
				strconv.FormatFloat(tx.Amount(), 'f', 2, 64),
				tx.Name(),
				tx.Category(),
			}, opts.Separator))
		}

		// skip credits
		if a < 0 {
			continue
		}

		// calculate paid
		if a < credit {
			paid += a
		}

		// sub debits
		credit -= a
	}

	sep := opts.Separator
	fmt.Fprint(tw, strings.Join([]string{
		"",
		fmt.Sprintf("total:%s%.2f%s((at least partially) unpaid and payments)", sep, total, sep),
		fmt.Sprintf("paid:%s%.2f%s(completely covered by payments)", sep, paid, sep),
		fmt.Sprintf("balance:%s%.2f", sep, credit),
		"",
	}, "\n"))

	tw.Flush()
}
