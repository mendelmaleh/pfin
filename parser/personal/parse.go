package personal

import (
	"bytes"
	"encoding/csv"
	"strings"

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

func (Parser) Parse(acc pfin.Account, filename string, data []byte) (txns []pfin.Transaction, err error) {
	raw, err := Parse(data)
	if err != nil {
		return txns, err
	}

	txns = make([]pfin.Transaction, len(raw))
	for i, v := range raw {
		before, _, _ := strings.Cut(filename, ".")
		v.Fields.User = before
		v.Fields.Account = acc.Name

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
