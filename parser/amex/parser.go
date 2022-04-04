package amex

import (
	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func Parse(data []byte, txns *[]pfin.Transaction) error {
	var raw []RawTransaction
	if err := csvutil.Unmarshal(data, &raw); err != nil {
		return err
	}

	// reverse order so it's chronological
	for i := len(raw) - 1; i != 0; i-- {
		*txns = append(*txns, raw[i])
	}

	return nil
}
