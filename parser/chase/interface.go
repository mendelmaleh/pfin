package chase

import (
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
	"git.sr.ht/~mendelmaleh/pfin/util"
)

var _ = pfin.Transaction(Transaction{})

func (tx Transaction) Date() time.Time {
	return tx.Raw.TransactionDate.Time
}

func (tx Transaction) Amount() float64 {
	return -tx.Raw.Amount
}

func (tx Transaction) Name() string {
	return tx.Raw.Description
}

func (tx Transaction) Category() string {
	return tx.Raw.Category
}

// not implemented
func (tx Transaction) Card() string {
	return ""
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
