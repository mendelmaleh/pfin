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

func (Parser) Parse(acc pfin.Account, data []byte) (txns []pfin.Transaction, err error) {
	var raw []RawTransaction
	if err = csvutil.Unmarshal(data, &raw); err != nil {
		return
	}

	length := len(raw)
	txns = make([]pfin.Transaction, length)

	// reverse order so it's chronological
	for i := 0; i < length; i++ {
		v := raw[length-i-1]

		v.UserField = acc.User(v.Card())
		txns[i] = v
	}

	return
}
