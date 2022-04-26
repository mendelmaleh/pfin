package bofa

import "git.sr.ht/~mendelmaleh/pfin/parser/util"

type Transaction struct {
	Fields // computed fields
	Raw    // csv fields
}

type Fields struct {
	User    string
	Account string
}

type Raw struct {
	Date           util.DateUS `csv:"Date"`
	Description    string      `csv:"Description"`
	Amount         float64     `csv:"Amount,omitempty"`
	RunningBalance float64     `csv:"Running Bal."`
}
