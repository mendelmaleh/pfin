package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	r := csv.NewReader(os.Stdin)

	// get header
	header, err := r.Read()
	if err != nil {
		panic(err)
	}

	// get col position
	var pos int = -1

	for i, v := range header {
		if v == "Credit" {
			pos = i
			break
		}
	}

	if pos == -1 {
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
