package util

import (
	"os"
	"path/filepath"
	"sort"

	"git.sr.ht/~mendelmaleh/pfin"
)

func ParseDir(parser, path string) ([]pfin.Transaction, error) {
	var txns []pfin.Transaction

	filetype, err := pfin.Filetype(parser)
	if err != nil {
		return txns, err
	}

	if path[len(path)-1] != filepath.Separator {
		path += string(filepath.Separator)
	}

	matches, err := filepath.Glob(path + "*." + filetype)
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

		tx, err := pfin.Parse(parser, file)
		if err != nil {
			return txns, err
		}

		// TODO: use copy(), consider parallelizing
		txns = append(txns, tx...)
	}

	return txns, nil
}
