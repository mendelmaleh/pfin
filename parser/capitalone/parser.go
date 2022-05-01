package capitalone

import (
	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func init() {
	pfin.Register("capitalone", Parser{})
}

// implements pfin.Parser interface
type Parser struct{}

func (Parser) Filetype() string {
	return "csv"
}

func (Parser) Parse(acc pfin.Account, filename string, data []byte) (txns []pfin.Transaction, err error) {
	var raw []Transaction
	if err = csvutil.Unmarshal(data, &raw); err != nil {
		return
	}

	length := len(raw)
	txns = make([]pfin.Transaction, length)

	// reverse order so it's chronological
	for i := 0; i < length; i++ {
		v := raw[length-i-1]

		v.Fields.User = acc.User(v.Card())
		v.Fields.Account = acc.Name

		txns[i] = v
	}

	return
}
