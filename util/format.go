package util

import (
	"strconv"
	"strings"
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
)

const ISO8601 = "2006-01-02"

func FormatCents(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func FormatDate(t time.Time) string {
	return t.Format(ISO8601)
}

func FormatFields(data [][]string, sep string) string {
	rows := make([]string, len(data)+1)
	for i, row := range data {
		rows[i] = strings.Join(row, sep)
	}

	return strings.Join(rows, "\n")
}

func FormatTx(tx pfin.Transaction, sep string) string {
	return strings.Join([]string{
		FormatDate(tx.Date()),
		FormatCents(tx.Amount()),
		tx.Account(),
		tx.Card(),
		tx.User(),
		tx.Name(),
		tx.Category(),
	}, sep)
}
