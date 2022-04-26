package personal

import (
	"bytes"
	"encoding/csv"

	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func init() {
	pfin.Register("personal", Parser{})
}

type Parser struct{}

func (Parser) Filetype() string {
	return "tsv"
}

func (Parser) Parse(acc pfin.Account, data []byte) (txns []pfin.Transaction, err error) {
	raw, err := Parse(data)
	if err != nil {
		return txns, err
	}

	txns = make([]pfin.Transaction, len(raw))
	for i, v := range raw {
		// there is no card/user data, so use default
		v.Fields.User = acc.User("")
		txns[i] = v
	}

	return
}

func Parse(data []byte) (raw []Transaction, err error) {
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = '\t'
	r.Comment = '-'

	dec, err := csvutil.NewDecoder(r)
	if err != nil {
		return raw, err
	}

	err = dec.Decode(&raw)
	return
}
