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
	Parse(data []byte) ([]Transaction, error)
	Filetype() string
}

func Register(name string, parser Parser) {
	parsers[name] = parser
}

func Parse(parser string, data []byte) ([]Transaction, error) {
	if _, ok := parsers[parser]; !ok {
		return []Transaction{}, ErrUnregisteredParser{parser}
	}

	return parsers[parser].Parse(data)
}

func Filetype(parser string) (string, error) {
	if _, ok := parsers[parser]; !ok {
		return "", ErrUnregisteredParser{parser}
	}

	return parsers[parser].Filetype(), nil
}
