package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"git.sr.ht/~mendelmaleh/pfin/parser/capitalone"
	"git.sr.ht/~mendelmaleh/pfin/parser/util"
	"github.com/jszwec/csvutil"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := os.ReadFile("sample.json")
	check(err)

	var raw T
	err = json.Unmarshal(data, &raw)
	check(err)

	last, err := util.NewDateISO("2023-12-15")
	check(err)

	var txns []capitalone.Raw
	for _, v := range raw.Entries {
		if t := v.TransactionBeginningStatementTimeStamp; t.IsZero() || t.After(last.Time) {
			tx := capitalone.Raw{
				TransactionDate: util.DateISO{v.TransactionDate},
				PostedDate:      util.DateISO{v.TransactionPostedDate},
				CardNumber:      v.TransactingCardLastFour,
				Description:     v.TransactionDescription,
				Category:        v.DisplayCategory,
				Debit:           0,
				Credit:          0,
			}

			if tx.TransactionDate.IsZero() {
				tx.TransactionDate = util.DateISO{v.TransactionDisplayDate}
			}

			if s := v.StatementDescription; s != "" {
				tx.Description = strings.TrimSpace(strings.Split(s, v.TransactionMerchant.Address.City)[0])
			}

			switch v.TransactionDebitCredit {
			case "Debit":
				tx.Debit = v.TransactionAmount
			case "Credit":
				tx.Credit = v.TransactionAmount
			default:
				check(fmt.Errorf("unexpected tx type %q", v.TransactionDebitCredit))
			}

			txns = append(txns, tx)
		}
	}

	w := csv.NewWriter(os.Stdout)
	enc := csvutil.NewEncoder(w)
	enc.Register(func(f float64) ([]byte, error) {
		if f == 0 {
			return []byte{}, nil
		}
		return []byte(strconv.FormatFloat(f, 'f', 2, 64)), nil
	})

	err = enc.Encode(txns)
	check(err)
	w.Flush()
}
