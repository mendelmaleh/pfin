package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/trimmer-io/go-csv"
)

func main() {
	opts := struct {
		unpaid *bool
	}{
		flag.Bool("unpaid", false, "print unpaid transactions"),
	}

	flag.Parse()

	var rows []Row

	dec := csv.NewDecoder(os.Stdin)
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
	fmt.Fprintf(os.Stderr, "total %.2f, debits %.2f, credits %.2f, ", debits+credits, debits, credits)

	// unpaid
	if *opts.unpaid {
		var unpaid float64
		var i int

		enc := csv.NewEncoder(os.Stdout)

		// TODO: improve float comparison hack
		for i = len(rows) - 1; int(((debits+credits)-unpaid)*100) > 0; i-- {
			if a := rows[i].Amount(); a > 0 {
				unpaid += a
				enc.EncodeRecord(rows[i])
			}
		}

		fmt.Fprintf(os.Stderr, "unpaid %.2f\n", unpaid)
	}
}
