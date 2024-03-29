package pfin

import "fmt"

var parsers = make(map[string]Parser)

type ErrUnregisteredParser struct {
	parser string
}

func (e ErrUnregisteredParser) Error() string {
	return fmt.Sprintf("pfin: unregistered parser %q", e.parser)
}

type Parser interface {
	Filetype() string
	Parse(acc Account, filename string, data []byte) ([]Transaction, error)
}

func Register(name string, parser Parser) {
	parsers[name] = parser
}

func Parse(acc Account, filename string, data []byte) ([]Transaction, error) {
	if _, ok := parsers[acc.Type]; !ok {
		return []Transaction{}, ErrUnregisteredParser{acc.Type}
	}

	txns, err := parsers[acc.Type].Parse(acc, filename, data)
	if err != nil {
		return []Transaction{}, fmt.Errorf("error parsing %s/%s: %w", acc.Name, filename, err)
	}

	return txns, nil
}

func Filetype(parser string) (string, error) {
	if _, ok := parsers[parser]; !ok {
		return "", ErrUnregisteredParser{parser}
	}

	return parsers[parser].Filetype(), nil
}
