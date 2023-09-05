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
	Users, Accounts util.StringFilter
}

func main() {
	// tool flags
	var opts Opts
	flag.StringVar(&opts.Users.String, "user", "", "filter user")
	flag.StringVar(&opts.Accounts.String, "account", "", "filter account")

	flag.Parse()

	// pfin config
	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal("error parsing config: ", err)
	}

	// collect accounts
	var txns []pfin.Transaction
	for name, acc := range config.Account {
		if opts.Accounts.Filter(name) {
			continue
		}

		tx, err := util.ParseDir(acc, config.Pfin.Root)
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

	users := map[string]int{}
	sums := []struct{ debits, credits float64 }{}

	var debits, credits float64
	for _, tx := range txns {
		u := tx.User()
		a := tx.Amount()

		if opts.Users.Filter(u) {
			continue
		}

		if _, ok := users[u]; !ok {
			users[u] = len(users)
			sums = append(sums, struct{ debits, credits float64 }{})
		}

		if a > 0 {
			debits += a
			sums[users[u]].debits += a
		} else {
			credits += a
			sums[users[u]].credits += a
		}

		fmt.Fprintln(tw, strings.Join([]string{
			util.FormatDate(tx.Date()),
			util.FormatCents(tx.Amount()),
			tx.Card(),
			tx.User(),
			tx.Name(),
			// tx.Category(),
		}, "\t"))

	}

	tw.Flush()
	fmt.Println()

	tw.Init(os.Stdout, 0, 8, 1, '\t', 1)

	fmt.Fprintf(tw, "name\tdebits\tcredits\ttotal\n")
	fmt.Fprintf(tw, "----\t------\t-------\t-----\n")
	fmt.Fprintf(tw, "%s\t%.2f\t%.2f\t%.2f\n", "total", debits, credits, debits+credits)

	for k, v := range users {
		fmt.Fprintf(tw, "%s\t%.2f\t%.2f\t%.2f\n", k, sums[v].debits, sums[v].credits, sums[v].debits+sums[v].credits)
	}

	tw.Flush()
}
