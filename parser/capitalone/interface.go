package capitalone

import (
	"fmt"
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
)

func (tx Transaction) Date() time.Time {
	return tx.Raw.TransactionDate.Time
}

func (tx Transaction) Amount() float64 {
	if tx.Debit != 0 {
		return tx.Debit
	}

	if tx.Credit != 0 {
		return -tx.Credit
	}

	// TODO
	panic("both credit and debit are zero: " + fmt.Sprintln(tx))

	return tx.Debit
}

func (tx Transaction) Name() string {
	return tx.Raw.Description
}

func (tx Transaction) Category() string {
	return tx.Raw.Category
}

func (tx Transaction) Card() string {
	return tx.Raw.CardNumber
}

func (tx Transaction) User() string {
	return tx.Fields.User
}

func (tx Transaction) Account() string {
	return tx.Fields.Account
}

func (tx Transaction) String() string {
	return pfin.TxString(tx, " ")
}
