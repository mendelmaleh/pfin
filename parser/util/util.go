package util

import (
	"os"
	"path/filepath"
	"sort"

	"git.sr.ht/~mendelmaleh/pfin"
)

func ParseDir(path string, parsefn func([]byte, *[]pfin.Transaction) error) ([]pfin.Transaction, error) {
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

		if err := parsefn(file, &txns); err != nil {
			return txns, err
		}
	}

	return txns, nil
}
