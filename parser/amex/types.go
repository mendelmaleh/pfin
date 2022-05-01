package amex

import (
	"fmt"
	"strings"

	"git.sr.ht/~mendelmaleh/pfin/parser/util"
)

type Transaction struct {
	Fields // computed fields
	Raw    // csv fields
}

type Fields struct {
	User    string
	Account string
}

type Raw struct {
	Date                     util.DateUS `csv:"Date"`
	Description              string      `csv:"Description"`
	CardMember               string      `csv:"Card Member"`
	AccountNumber            string      `csv:"Account #"`
	Amount                   float64     `csv:"Amount"`
	ExtendedDetails          string      `csv:"Extended Details"`
	AppearsOnYourStatementAs string      `csv:"Appears On Your Statement As"`
	Address                  string      `csv:"Address"`

	CityState `csv:"City/State"`

	ZipCode   string `csv:"Zip Code"`
	Country   string `csv:"Country"`
	Reference string `csv:"Reference"`

	Category `csv:"Category"`
}

type CityState struct {
	City  string
	State string
}

func (c *CityState) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	s := strings.Split(string(data), "\n")

	if len(s) >= 1 {
		c.City = s[0]
	}

	if len(s) == 2 {
		c.State = s[1]
	}

	if len(s) > 2 {
		return fmt.Errorf("expected two parts for City/State, got %d (%v, %s)", len(s), s, string(data))
	}

	return nil
}

type Category struct {
	Main string
	Sub  string
}

func (c *Category) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	s := strings.Split(string(data), "-")

	if len(s) >= 1 {
		c.Main = s[0]
	}

	if len(s) == 2 {
		c.Sub = s[1]
	}

	if len(s) > 2 {
		return fmt.Errorf("expected two parts for Category, got %d (%v, %s)", len(s), s, string(data))
	}

	return nil
}

func (c *Category) String() string {
	return fmt.Sprintf("%s/%s", c.Main, c.Sub)
}
