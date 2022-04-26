package capitalone

import "git.sr.ht/~mendelmaleh/pfin/parser/util"

type Transaction struct {
	Fields // computed fields, namespaced so they don't conflict with the interface methods
	Raw    // raw fields from the csv
}

type Fields struct {
	User string
}

type Raw struct {
	TransactionDate util.Date `csv:"Transaction Date"`
	PostedDate      util.Date `csv:"Posted Date"`
	CardNumber      string    `csv:"Card No."`
	Description     string    `csv:"Description"`
	Category        string    `csv:"Category"`
	Debit           float64   `csv:"Debit,omitempty"`
	Credit          float64   `csv:"Credit,omitempty"`
}
