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

func TxString(tx pfin.Transaction, sep string) string {
	var b strings.Builder

	b.WriteString(FormatDate(tx.Date()))
	b.WriteString(sep)

	b.WriteString(FormatCents(tx.Amount()))
	b.WriteString(sep)

	b.WriteString(tx.Account())
	b.WriteString(sep)

	b.WriteString(tx.User())
	b.WriteString(sep)

	b.WriteString(tx.Name())
	b.WriteString(sep)

	b.WriteString(tx.Category())

	return b.String()
}
