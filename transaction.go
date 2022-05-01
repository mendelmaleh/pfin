package pfin

import (
	"strconv"
	"strings"
	"time"
)

const ISO8601 = "2006-01-02"

type Transaction interface {
	Date() time.Time
	Amount() float64
	Name() string

	Category() string

	Card() string // may be empty
	User() string
	Account() string

	// should be pfin.TxString
	String() string
}

func TxString(tx Transaction, sep string) string {
	var b strings.Builder

	b.WriteString(tx.Date().Format(ISO8601))
	b.WriteString(sep)

	b.WriteString(strconv.FormatFloat(tx.Amount(), 'f', 2, 64))
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
