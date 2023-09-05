package citi

import (
	"bytes"
	"regexp"

	"git.sr.ht/~mendelmaleh/pfin"
	"golang.org/x/net/html"
)

func init() {
	pfin.Register("citi", Parser{})
}

var (
	digitalRegex = regexp.MustCompile(` Digital Account Number XXXXXXXXXXXX(\d{4})$`)
	virtualRegex = regexp.MustCompile(` - Virtual Account Number (\d{4})$`)
)

// implements pfin.Parser interface
type Parser struct{}

func (Parser) Filetype() string {
	return "html"
}

func (Parser) Parse(acc pfin.Account, filename string, data []byte) (txns []pfin.Transaction, err error) {
	// TODO: user io.Reader interface in parser.Parse
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		// TODO: test err shadowing
		return
	}

	for _, v := range SelectorTxRow.MatchAll(doc) {
		tx, err := TxFromHTML(v)
		if err != nil {
			return txns, err
		}

		// virtual card number or cardmember
		if matches := virtualRegex.FindStringSubmatchIndex(tx.Description); len(matches) > 0 {
			tx.Fields.Card = tx.Description[matches[2]:]
			tx.Description = tx.Description[:matches[0]]
		} else if matches := digitalRegex.FindStringSubmatchIndex(tx.Description); len(matches) > 0 {
			tx.Fields.Card = tx.Description[matches[2]:]
			tx.Description = tx.Description[:matches[0]]
		} else {
			tx.Fields.Card = tx.Raw.Cardmember
		}

		tx.Fields.User = acc.Cards[tx.Fields.Card]
		tx.Fields.Account = acc.Name

		txns = append(txns, tx)
	}

	return
}
