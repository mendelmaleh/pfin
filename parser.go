package pfin

import (
	"errors"
)

var parsers = make(map[string]Parser)

var ErrUnregisteredParser = errors.New("pfin: unregistered parser")

type Parser interface {
	Parse(data []byte) ([]Transaction, error)
	Filetype() string
}

func Register(name string, parser Parser) {
	parsers[name] = parser
}

func Parse(parser string, data []byte) ([]Transaction, error) {
	if _, ok := parsers[parser]; !ok {
		return []Transaction{}, ErrUnregisteredParser
	}

	return parsers[parser].Parse(data)
}

func Filetype(parser string) (string, error) {
	if _, ok := parsers[parser]; !ok {
		return "", ErrUnregisteredParser
	}

	return parsers[parser].Filetype(), nil
}
