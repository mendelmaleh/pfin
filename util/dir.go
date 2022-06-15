package util

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"git.sr.ht/~mendelmaleh/pfin"
)

type ErrNoMatches struct {
	path string
}

func (e ErrNoMatches) Error() string {
	return fmt.Sprintf("pfin/parser/util: no matches for path %q", e.path)
}

func MatchDir(acc pfin.Account, root string) ([]string, error) {
	var matches []string

	filetype, err := pfin.Filetype(acc.Type)
	if err != nil {
		return matches, err
	}

	path := filepath.Join(root, acc.Name)

	if path[len(path)-1] != filepath.Separator {
		path += string(filepath.Separator)
	}

	matches, err = filepath.Glob(path + "*." + filetype)
	if err != nil {
		return matches, err
	}

	if len(matches) == 0 {
		return matches, ErrNoMatches{path}
	}

	// sort oldest first
	sort.Strings(matches)

	return matches, nil
}

func ParseDir(acc pfin.Account, root string) ([]pfin.Transaction, error) {
	var txns []pfin.Transaction

	matches, err := MatchDir(acc, root)
	if err != nil {
		return txns, err
	}

	for _, f := range matches {
		file, err := os.ReadFile(f)
		if err != nil {
			return txns, err
		}

		tx, err := pfin.Parse(acc, filepath.Base(f), file)
		if err != nil {
			return txns, err
		}

		// TODO: use copy(), consider parallelizing
		txns = append(txns, tx...)
	}

	return txns, nil
}
