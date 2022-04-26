package bofa

import (
	"time"

	"git.sr.ht/~mendelmaleh/pfin"
)

func (tx RawTransaction) Date() time.Time {
	return tx.Fields.Date.Time
}

func (tx RawTransaction) Amount() float64 {
	return tx.Fields.Amount
}

func (tx RawTransaction) Name() string {
	return tx.Description
}

func (tx RawTransaction) Card() string {
	return ""
}

func (tx RawTransaction) User() string {
	return tx.UserField
}

func (tx RawTransaction) String() string {
	return pfin.TxString(tx, " ")
}
