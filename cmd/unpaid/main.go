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
				tx.Date().Format(pfin.ISO8601),
				strconv.FormatFloat(tx.Amount(), 'f', 2, 64),
				tx.Name(),
				tx.Category(),
			}, opts.Separator))
		}

		// update balance
		balance -= a
	}

	sep := opts.Separator
	fmt.Fprint(tw, strings.Join([]string{
		"",
		fmt.Sprintf("Total:%s%.2f%s((at least partially) unpaid)", sep, total, sep),
		fmt.Sprintf("Paid:%s%.2f%s(completely covered by payments)", sep, paid, sep),
		fmt.Sprintf("Payments:%s%.2f%s(previous payments)", sep, credit, sep),
		fmt.Sprintf("Balance:%s%.2f%s(unpaid + paid - payments) * -1", sep, balance, sep),
		"",
		"By category:",
		"",
	}, "\n"))

	var sorted []string
	for k, _ := range categories {
		sorted = append(sorted, k)
	}

	sort.Strings(sorted)
	for _, k := range sorted {
		v := categories[k]
		fmt.Fprintf(tw, "%s:%s%.2f\n", k, sep, v)
	}

	tw.Flush()
}
