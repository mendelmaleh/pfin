package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/all"
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

	var debits []pfin.Transaction
	var credit float64

	for name, acc := range config.Account {
		if opts.Accounts.Filter(name) && opts.Payments.Filter(name) {
			continue
		}

		txns, err := util.ParseDir(acc, config.Pfin.Root)
		if err != nil {
			log.Fatal(err)
		}

		for _, tx := range txns {
			if opts.Users.Filter(tx.User()) {
				continue
			}

			if tx.Amount() < 0 {
				credit += tx.Amount()
			} else {
				debits = append(debits, tx)
			}

		}
	}

	sort.SliceStable(debits, func(i, j int) bool {
		return debits[i].Date().Before(debits[j].Date())
	})

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	categories := make(map[string]float64)
	balance := credit * -1

	var total, paid float64
	for _, tx := range debits {
		a := tx.Amount()

		// calculate paid
		if a < balance {
			paid += a
		}

		// print uncovered debits
		if a > balance {
			total += a

			// calculate unpaid categories
			c := tx.Category()

			if _, ok := categories[c]; !ok {
				categories[c] = a
			} else {
				categories[c] += a
			}

			fmt.Fprintln(tw, strings.Join([]string{
				util.FormatDate(tx.Date()),
				util.FormatCents(a),
				tx.Name(),
				tx.Category(),
			}, opts.Separator))
		}

		// update balance
		balance -= a
	}

	data := [][]string{
		{},
		{"Total:", util.FormatCents(total), "((at least partially) unpaid)"},
		{"Paid:", util.FormatCents(paid), "(completely covered by payments)"},
		{"Payments:", util.FormatCents(credit), "(previous payments)"},
		{"Balance:", util.FormatCents(balance), "(unpaid + paid - payments) * -1"},
		{},
		{"By category:"},
		{},
	}

	var sorted []string
	for k, _ := range categories {
		sorted = append(sorted, k)
	}

	sort.Strings(sorted)
	for _, k := range sorted {
		data = append(data, []string{k + ":", util.FormatCents(categories[k])})
	}

	fmt.Fprint(tw, util.FormatFields(data, opts.Separator))
	tw.Flush()
}
