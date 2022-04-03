package capitalone

import (
	"os"
	"path/filepath"
	"sort"

	"git.sr.ht/~mendelmaleh/pfin"
	"github.com/jszwec/csvutil"
)

func ParseDir(path string) ([]pfin.Transaction, error) {
	var txns []pfin.Transaction

	matches, err := filepath.Glob(path + "*.csv")
	if err != nil {
		return txns, err
	}

	// sort oldest first
	sort.Strings(matches)

	for _, f := range matches {
		file, err := os.ReadFile(f)
		if err != nil {
			return txns, err
		}

		var raw []RawTransaction
		if err := csvutil.Unmarshal(file, &raw); err != nil {
			return txns, err
		}

		// https://github.com/golang/go/wiki/SliceTricks#extend-capacity
		/*
			if cap(txns) - len(txns) < len(raw) {
				txns = append(make([]pfin.Transaction, 0, len(txns) + len(raw)), txns...)
			}
		*/

		// https://go.dev/doc/faq#convert_slice_of_interface
		// append in reverse
		for i := len(raw) - 1; i != 0; i-- {
			txns = append(txns, raw[i])
		}
	}

	return txns, nil
}
