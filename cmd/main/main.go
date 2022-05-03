package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"text/tabwriter"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/amex"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/bofa"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/capitalone"
	_ "git.sr.ht/~mendelmaleh/pfin/parser/personal"
)

type Opts struct {
	Users, Accounts util.StringFilter
}

func main() {
	// flags
	var opts Opts
	flag.StringVar(&opts.Users.String, "user", "", "filter user")
	flag.StringVar(&opts.Accounts.String, "account", "", "filter account")

	flag.Parse()

	// config
	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal("error parsing config: ", err)
	}

	// collect
	ch := make(chan []pfin.Transaction, len(config.Account))
	var wg sync.WaitGroup

	for name, acc := range config.Account {
		if opts.Accounts.Filter(name) {
			continue
		}

		wg.Add(1)
		go func(acc pfin.Account, root string) {
			defer wg.Done()
			tx, err := util.ParseDir(acc, root)
			if err != nil {
				panic(err)
			}

			ch <- tx
		}(acc, config.Pfin.Root)
	}

	wg.Wait()
	close(ch)

	var txns []pfin.Transaction
	for tx := range ch {
		txns = append(txns, tx...)
	}

	// sort oldest to newest
	sort.SliceStable(txns, func(i, j int) bool {
		return txns[i].Date().Before(txns[j].Date())
	})

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	var sum = map[string]float64{}

	for _, tx := range txns {
		if opts.Users.Filter(tx.User()) {
			continue
		}

		if _, ok := sum[tx.User()]; !ok {
			sum[tx.User()] = 0
		}

		sum[tx.User()] += tx.Amount()
		fmt.Fprintln(tw, util.FormatTx(tx, "\t"))
	}

	tw.Flush()
	fmt.Println()

	tw.Init(os.Stdout, 0, 8, 1, '\t', 0)
	for user, total := range sum {
		fmt.Fprintf(tw, "%s\t%.2f\n", user, total)
	}

	tw.Flush()
}
