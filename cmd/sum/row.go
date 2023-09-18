package main

import (
	"time"

	dates "git.sr.ht/~mendelmaleh/pfin/parser/util"
	"git.sr.ht/~mendelmaleh/pfin/util"
)

type Row struct {
	Fields
}

type Fields struct {
	Date     dates.DateISO `csv:"date"`
	Amount   float64       `csv:"amount"`
	Account  string        `csv:"account"`
	Card     string        `csv:"card"`
	User     string        `csv:"user"`
	Name     string        `csv:"name"`
	Category string        `csv:"category"`
}

func (r *Row) Date() time.Time {
	return r.Fields.Date.Time
}

func (r *Row) Amount() float64 {
	return r.Fields.Amount
}

func (r *Row) Name() string {
	return r.Fields.Name
}

func (r *Row) Category() string {
	return r.Fields.Category
}

func (r *Row) Card() string {
	return r.Fields.Card
}

func (r *Row) User() string {
	return r.Fields.User
}

func (r *Row) Account() string {
	return r.Fields.Account
}

// should be util.FormatTx
func (r *Row) String() string {
	return util.FormatTx(r, " ")
}
