package checking

import (
	"time"

	"git.sr.ht/~mendelmaleh/pfin/util"
)

func (tx Transaction) Date() time.Time {
	return tx.Raw.Date.Time
}

func (tx Transaction) Amount() float64 {
	return -tx.Raw.Amount.Amount
}

func (tx Transaction) Name() string {
	return tx.Raw.Description // not very descriptive
}

func (tx Transaction) Category() string {
	// not really category, more like transaction type
	return tx.Raw.Typename
}

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
