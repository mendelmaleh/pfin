package bofa

import (
	"bytes"

	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func Parse(data []byte, txns *[]pfin.Transaction) error {
	// bofa csv statements have two "tables" in them, the first is a summary, the second is the transactions.
	parts := bytes.Split(data, []byte("\r\n\r\n"))
	data = parts[1]

	var raw []RawTransaction
	if err := csvutil.Unmarshal(data, &raw); err != nil {
		return err
	}

	for _, tx := range raw {
		*txns = append(*txns, tx)
	}

	return nil
}
