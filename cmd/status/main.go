package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/parser/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/capitalone"
)

func main() {
	// parse config
	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	// parse accounts
	txns, err := util.ParseDir("capitalone", filepath.Join(config.Pfin.Root, "capitalone"))
	if err != nil {
		log.Fatal(err)
	}

	sort.SliceStable(txns, func(i, j int) bool {
		return txns[i].Date().Before(txns[j].Date())
	})

	// summary
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	for _, tx := range txns {
		fmt.Fprintln(tw, pfin.TxString(tx, "\t"))
	}

	tw.Flush()

	// days since last transaction
	last := txns[len(txns)-1]
	fmt.Println(last.Date())
	fmt.Println(int(time.Now().Sub(last.Date()).Hours()) % 24)
}
