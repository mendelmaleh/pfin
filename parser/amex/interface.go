package amex

import (
	"time"

	"git.sr.ht/~mendelmaleh/pfin/util"
)

func (tx Transaction) Date() time.Time {
	return tx.Raw.Date.Time
}

func (tx Transaction) Amount() float64 {
	return tx.Raw.Amount
}

func (tx Transaction) Name() string {
	return tx.Raw.Description
}

func (tx Transaction) Category() string {
	return tx.Raw.Category.String()
}

func (tx Transaction) Card() string {
	return tx.Raw.AccountNumber
}

func (tx Transaction) User() string {
	return tx.Fields.User
}

func (tx Transaction) Account() string {
	return tx.Fields.Account
}

func (tx Transaction) String() string {
	return util.FormatTx(tx, " ")
}
