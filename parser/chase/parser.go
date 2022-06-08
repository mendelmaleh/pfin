package chase

import (
	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func init() {
	pfin.Register("chase", Parser{})
}

type Parser struct{}

func (Parser) Parse(acc pfin.Account, filename string, data []byte) (txns []pfin.Transaction, err error) {
	var raw []Transaction
	if err = csvutil.Unmarshal(data, &raw); err != nil {
		return
	}

	length := len(raw)
	txns = make([]pfin.Transaction, length)

	// reverse order
	for i := 0; i < length; i++ {
		v := raw[length-i-1]

		// default user
		v.Fields.User = acc.User("")
		v.Fields.Account = acc.Name

		txns[i] = v
	}

	return
}

func (Parser) Filetype() string {
	return "csv"
}
