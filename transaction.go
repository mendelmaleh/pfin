package pfin

import (
	"time"
)

type Transaction interface {
	Date() time.Time
	Amount() float64
	Name() string

	Category() string

	Card() string // may be empty
	User() string
	Account() string

	// should be util.TxString
	String() string
}
