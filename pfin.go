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
	Card() string
	Name() string

	// should be pfin.TxString
	String() string
}

func TxString(tx Transaction, sep string) string {
	var b strings.Builder

	b.WriteString(tx.Date().Format(ISO8601))
	b.WriteString(sep)

	b.WriteString(strconv.FormatFloat(tx.Amount(), 'f', 2, 64))
	b.WriteString(sep)

	b.WriteString(tx.Card())
	b.WriteString(sep)

	b.WriteString(tx.Name())

	return b.String()
}
