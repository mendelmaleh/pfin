package bofa

import (
	"bytes"

	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func init() {
	pfin.Register("bofa", Parser{})
}

type Parser struct{}

func (Parser) Filetype() string {
	return "csv"
}

func (Parser) Parse(acc pfin.Account, filename string, data []byte) (txns []pfin.Transaction, err error) {
	// bofa csv statements have two "tables" in them, the first is a summary, the second is the transactions.
	parts := bytes.Split(data, []byte("\r\n\r\n"))
	data = parts[1]

	var raw []Transaction
	if err = csvutil.Unmarshal(data, &raw); err != nil {
		return
	}

	txns = make([]pfin.Transaction, len(raw))
	for i, v := range raw {
		v.Fields.User = acc.User(v.Card())
		v.Fields.Account = acc.Name

		txns[i] = v
	}

	return
}
