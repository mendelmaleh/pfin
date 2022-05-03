package util

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"git.sr.ht/~mendelmaleh/pfin"
)

type result struct {
	s  string
	tx []pfin.Transaction
}

type ErrNoMatches struct {
	path string
}

func (e ErrNoMatches) Error() string {
	return fmt.Sprintf("pfin/parser/util: no matches for path %q", e.path)
}

func ParseDir(acc pfin.Account, root string) ([]pfin.Transaction, error) {
	var txns []pfin.Transaction

	filetype, err := pfin.Filetype(acc.Type)
	if err != nil {
		return txns, err
	}

	path := filepath.Join(root, acc.Name)

	if path[len(path)-1] != filepath.Separator {
		path += string(filepath.Separator)
	}

	matches, err := filepath.Glob(path + "*." + filetype)
	if err != nil {
		return txns, err
	}

	if len(matches) == 0 {
		return txns, ErrNoMatches{path}
	}

	ch := make(chan result, len(matches))
	cherr := make(chan error, 1)

	var wg sync.WaitGroup

	for _, f := range matches {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()

			file, err := os.ReadFile(f)
			if err != nil {
				cherr <- err
				return
			}

			tx, err := pfin.Parse(acc, filepath.Base(f), file)
			if err != nil {
				cherr <- err
			}

			ch <- result{f, tx}
		}(f)
	}

	wg.Wait()
	close(ch)

	files := make(map[string][]pfin.Transaction)

	for len(files) < len(matches) {
		select {
		case err := <-cherr:
			return txns, err
		case res := <-ch:
			files[res.s] = res.tx
			// txns = append(txns, tx...)
		}
	}

	// sort oldest first
	sort.Strings(matches)
	for _, f := range matches {
		txns = append(txns, files[f]...)
	}

	return txns, nil
}
