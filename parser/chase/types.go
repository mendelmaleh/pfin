package chase

import "git.sr.ht/~mendelmaleh/pfin/parser/util"

type Transaction struct {
	Fields
	Raw
}

type Fields struct {
	User    string
	Account string
}

type Raw struct {
	TransactionDate util.DateUS `csv:"Transaction Date"`
	PostDate        util.DateUS `csv:"Post Date"`
	Description     string      `csv:"Description"`
	Category        string      `csv:"Category"`
	Type            string      `csv:"Type"`
	Amount          float64     `csv:"Amount"`
	Memo            string      `csv:"Memo"`
}
