package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"

	_ "git.sr.ht/~mendelmaleh/pfin/parser/all"
)

func header(cols []string) string {
	hdr := strings.Join(cols, "\t")
	div := strings.Map(func(r rune) rune {
		if r != '\t' {
			return '-'
		}
		return r
	}, hdr)

	return strings.Join([]string{hdr, div, ""}, "\n")
}

func main() {
	// opts
	var full bool
	flag.BoolVar(&full, "v", false, "verbose output (total, debits, credits)")
	flag.Parse()

	cols := []string{
		"statement",
		"balance",
		"total",
		"debits",
		"credits",
	}

	if !full {
		cols = cols[0:2]
	}

	// parse config
	config, err := pfin.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	// account
	name := flag.Arg(0)
	acc, ok := config.Account[name]
	if !ok {
		log.Fatalf("no account %q\n", name)
	}

	// parse dir
	matches, err := util.MatchDir(acc, config.Pfin.Root)
	if err != nil {
		log.Fatal(err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	defer tw.Flush()
	fmt.Fprint(tw, header(cols))

	// parse statements
	var balance float64
	for _, stmt := range matches {
		data, err := os.ReadFile(stmt)
		if err != nil {
			log.Fatal(err)
		}

		name := filepath.Base(stmt)
		name = name[0 : len(name)-len(filepath.Ext(name))]

		txns, err := pfin.Parse(acc, name, data)
		if err != nil {
			log.Fatal(err)
		}

		// statement totals
		var credits, debits float64
		for _, tx := range txns {
			if a := tx.Amount(); a < 0 {
				credits += a
			} else {
				debits += a
			}
		}

		total := credits + debits
		balance += total

		fmt.Fprint(tw, name)

		for _, v := range []float64{
			balance,
			total,
			debits,
			credits,
		} {
			fmt.Fprint(tw, "\t"+util.FormatCents(v))
			if !full {
				break // after balance
			}
		}

		fmt.Fprint(tw, "\n")
	}
}
