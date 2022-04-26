package personal

import "git.sr.ht/~mendelmaleh/pfin/parser/util"

type Transaction struct {
	Fields // computed fields
	Raw    // csv fields
}

type Fields struct {
	User string
}

type Raw struct {
	Date        util.DateISO `csv:"date    "`
	Amount      float64      `csv:"amount"`
	Media       string       `csv:"media"`
	Description string       `csv:"description"`
}
