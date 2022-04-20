package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type Opts struct {
	Field   string
	Comma   string
	Comment string
	Trim    bool
	Untab   bool

	PSV bool
	TSV bool
}

func main() {
	// flags
	var opts Opts

	flag.StringVar(&opts.Field, "field", "", "column to sum")
	flag.StringVar(&opts.Comma, "comma", ",", "column separator")
	flag.StringVar(&opts.Comment, "comment", "#", "column separator")
	flag.BoolVar(&opts.Trim, "trim", false, "trim leading space")
	flag.BoolVar(&opts.Untab, "untab", false, "remove tabs before parsing")

	// presets
	flag.BoolVar(&opts.PSV, "psv", false, "use psv preset")
	flag.BoolVar(&opts.TSV, "tsv", false, "use tsv preset")

	flag.Parse()

	if opts.PSV {
		opts.Field = "amount"
		opts.Comma = "|"
		opts.Comment = "-"
		opts.Trim = true
		opts.Untab = true
	}

	if opts.TSV {
		opts.Comma = "\t"
		opts.Comment = "-"
	}

	// data
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	// untab before parsing
	if opts.Untab {
		data = bytes.ReplaceAll(data, []byte{'\t'}, []byte{})
	}

	// csv parser
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = []rune(opts.Comma)[0]
	r.Comment = []rune(opts.Comment)[0]
	r.TrimLeadingSpace = opts.Trim

	// get header
	header, err := r.Read()
	if err != nil {
		panic(err)
	}

	// get col position
	var pos int = -1

	for i, v := range header {
		if v == opts.Field {
			pos = i
			break
		}
	}

	if pos == -1 {
		spew.Dump(opts)
		panic("can't find position in header")
	}

	// sum
	var sum float64

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		if rec[pos] != "" {
			f64, err := strconv.ParseFloat(rec[pos], 64)
			if err != nil {
				panic(err)
			}

			sum += f64
		}
	}

	fmt.Printf("%.2f\n", sum)
}
