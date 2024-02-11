package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jszwec/csvutil"
)

func main() {
	opts := struct {
		unpaid *bool
	}{
		flag.Bool("unpaid", false, "print unpaid transactions"),
	}

	flag.Parse()

	var rows []Row

	header, err := csvutil.Header(Row{}, "")
	if err != nil {
		log.Fatal(err)
	}

	dec, err := csvutil.NewDecoder(csv.NewReader(os.Stdin), header...)
	if err != nil {
		log.Fatal(err)
	}

	if err := dec.Decode(&rows); err != nil {
		log.Fatal(err)
	}

	/*
		data, _ := io.ReadAll(os.Stdin)
		// fmt.Printf("%s", data)

		if err := csvutil.Unmarshal(data, &rows); err != nil {
			log.Fatal(err)
		}
	*/

	/*
		dec, err := csvutil.NewDecoder(csv.NewReader(os.Stdin))
		if err != nil {
			log.Fatal(err)
		}

		if err := dec.Decode(&rows); err != nil {
			log.Fatal(err)
		}
	*/

	// fmt.Println(rows)

	var debits, credits float64

	for _, tx := range rows {
		a := tx.Amount()
		if a > 0 {
			debits += a
		} else {
			credits += a
		}
	}

	// summary
	fmt.Fprintf(os.Stderr, "total %.2f, debits %.2f, credits %.2f", debits+credits, debits, credits)

	writer := csv.NewWriter(os.Stdout)
	enc := csvutil.NewEncoder(writer)

	// unpaid
	if *opts.unpaid {
		var unpaid float64
		var i int

		// TODO: improve float comparison hack
		for i = len(rows) - 1; int(((debits+credits)-unpaid)*100) > 0; i-- {
			if a := rows[i].Amount(); a > 0 {
				unpaid += a
				enc.Encode(rows[i])
			}
		}

		fmt.Fprintf(os.Stderr, ", unpaid %.2f", unpaid)
	} else {
		for i := len(rows) - 1; i >= 0; i-- {
			enc.Encode(rows[i])
		}
	}

	writer.Flush()
	fmt.Fprintf(os.Stderr, "\n")
}
