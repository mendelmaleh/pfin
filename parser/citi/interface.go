package citi

import (
	"time"

	"git.sr.ht/~mendelmaleh/pfin/util"
)

func (tx Transaction) Date() time.Time {
	if tx.Purchased.IsZero() {
		return tx.Posted
	}

	return tx.Purchased
}

func (tx Transaction) Amount() float64 {
	return tx.Raw.Amount
}

func (tx Transaction) Name() string {
	return tx.Raw.Description
}

func (tx Transaction) Category() string {
	return tx.Raw.Category
}

func (tx Transaction) Card() string {
	return tx.Fields.Card
}

func (tx Transaction) User() string {
	return tx.Fields.User
}

func (tx Transaction) Account() string {
	return tx.Fields.Account
}

// should be util.FormatTx
func (tx Transaction) String() string {
	return util.FormatTx(tx, " ")
}
