package bofa

import (
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
)

func (tx RawTransaction) Date() time.Time {
	return tx.DateField.Time
}

func (tx RawTransaction) Amount() float64 {
	return tx.AmountField
}

func (tx RawTransaction) Card() string {
	return "-"
}

func (tx RawTransaction) Name() string {
	return tx.Description
}

func (tx RawTransaction) String() string {
	return pfin.TxString(tx, " ")
}
