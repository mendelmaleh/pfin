package checking

import (
	"encoding/json"

	"git.sr.ht/~mendelmaleh/pfin"
)

func init() {
	pfin.Register("amex-checking", Parser{})
}

type Parser struct{}

func (Parser) Filetype() string {
	return "json"
}

func (Parser) Parse(acc pfin.Account, filename string, data []byte) (txns []pfin.Transaction, err error) {
	var resp Response
	if err = json.Unmarshal(data, &resp); err != nil {
		return
	}

	raw := resp.Data.ProductAccountByAccountNumberProxy.FinancialPosition.Transactions.CheckingTransactions
	length := len(raw)
	txns = make([]pfin.Transaction, length)

	for i := 0; i < length; i++ {
		v := raw[length-i-1]

		v.Fields.Account = acc.Name
		v.Fields.User = acc.User(v.Card())

		txns[i] = v
	}

	return
}
