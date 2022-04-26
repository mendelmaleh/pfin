package personal

import "git.sr.ht/~mendelmaleh/pfin/parser/util"

type Fields struct {
	Date        util.Date `csv:"date    "`
	Amount      float64   `csv:"amount"`
	Media       string    `csv:"media"`
	Description string    `csv:"description"`
}

type RawTransaction struct {
	Fields
	UserField string
}
