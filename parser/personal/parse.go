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

func (Parser) Parse(data []byte) (txns []pfin.Transaction, err error) {
	// remove tabs
	data = bytes.ReplaceAll(data, []byte{'\t'}, []byte{})

	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = '\t'
	r.Comment = '-'

	dec, err := csvutil.NewDecoder(r)
	if err != nil {
		return
	}

	var raw []RawTransaction
	if err = dec.Decode(&raw); err != nil {
		return
	}

	txns = make([]pfin.Transaction, len(raw))
	for i, v := range raw {
		v.UserField = "-"
		txns[i] = v
	}

	return
}
