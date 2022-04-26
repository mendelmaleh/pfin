package amex

import (
	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func init() {
	pfin.Register("amex", Parser{})
}

type Parser struct{}

func (Parser) Filetype() string {
	return "csv"
}

func (Parser) Parse(data []byte) (txns []pfin.Transaction, err error) {
	var raw []RawTransaction
	if err = csvutil.Unmarshal(data, &raw); err != nil {
		return
	}

	length := len(raw)
	txns = make([]pfin.Transaction, length)

	// reverse order so it's chronological
	for i := 0; i < length; i++ {
		txns[i] = raw[length-i-1]
	}

	return
}
