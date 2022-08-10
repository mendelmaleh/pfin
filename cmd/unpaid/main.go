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
	var debit, credit float64

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

			if a := tx.Amount(); a > 0 {
				debits = append(debits, tx)
				debit += a
			} else {
				// credits = append(credits, tx)
				credit += a
			}
		}
	}

	sort.SliceStable(debits, func(i, j int) bool {
		return debits[i].Date().Before(debits[j].Date())
	})

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	i := 0

	var paid float64
	for i < len(debits) {
		a := debits[i].Amount()

		if paid+a > -credit {
			break
		}

		paid += a
		i++
	}

	var unpaid float64
	for ; i < len(debits); i++ {
		tx := debits[i]
		unpaid += tx.Amount()

		fmt.Fprintln(tw, strings.Join([]string{
			util.FormatDate(tx.Date()),
			util.FormatCents(tx.Amount()),
			tx.Name(),
			tx.Category(),
		}, opts.Separator))
	}

	fmt.Fprint(tw, "\n")
	tw.Flush()

	tw.Init(os.Stdout, 1, 8, 2, ' ', 0)

	func(keys ...string) {
		header := strings.Join(keys, "\t")
		fmt.Fprintln(tw, header)

		fmt.Fprintln(tw, strings.Map(func(r rune) rune {
			if r != '\t' {
				return '-'
			}
			return r
		}, header))
	}("balance", "unpaid", "paid", "debits", "credits")

	func(values ...float64) {
		var b strings.Builder

		for i, v := range values {
			if i != 0 {
				b.WriteString("\t")
			}

			b.WriteString(util.FormatCents(v))
		}

		fmt.Fprintln(tw, b.String())
	}(debit+credit, unpaid, paid, debit, credit)

	tw.Flush()
}
