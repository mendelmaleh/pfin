package capitalone

import "git.sr.ht/~mendelmaleh/pfin/parser/util"

type RawTransaction struct {
	TransactionDate util.Date `csv:"Transaction Date"`
	PostedDate      util.Date `csv:"Posted Date"`
	CardNumber      string    `csv:"Card No."`
	Description     string    `csv:"Description"`
	Category        string    `csv:"Category"`
	Debit           float64   `csv:"Debit,omitempty"`
	Credit          float64   `csv:"Credit,omitempty"`
}
