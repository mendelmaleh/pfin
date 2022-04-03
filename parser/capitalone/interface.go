package capitalone

import (
	"fmt"
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
)

func (tx RawTransaction) Date() time.Time {
	return tx.TransactionDate.Time
}

func (tx RawTransaction) Amount() float64 {
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

func (tx RawTransaction) Name() string {
	return tx.Description
}

func (tx RawTransaction) String() string {
	return pfin.TxString(tx, " ")
}
