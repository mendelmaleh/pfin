package main

import (
	"fmt"
	"log"
	"sort"
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

	// parse accounts
	txns := make(map[string][]pfin.Transaction, len(config.Account))

	for name, acc := range config.Account {
		tx, err := util.ParseDir(acc, config.Pfin.Root)
		if err != nil {
			log.Fatal(err)
		}

		sort.SliceStable(tx, func(i, j int) bool {
			return tx[i].Date().Before(tx[j].Date())
		})

		txns[name] = tx
	}

	// days since last transaction
	for name, tx := range txns {
		last := tx[len(tx)-1].Date()
		fmt.Printf("%s: %d\n", name, int(time.Now().Sub(last).Hours())/24)
	}
}
