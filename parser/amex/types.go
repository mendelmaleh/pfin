package amex

import (
	"fmt"
	"strings"
	"time"
)

type Fields struct {
	Date                     Date    `csv:"Date"`
	Description              string  `csv:"Description"`
	CardMember               string  `csv:"Card Member"`
	AccountNumber            string  `csv:"Account #"`
	Amount                   float64 `csv:"Amount"`
	ExtendedDetails          string  `csv:"Extended Details"`
	AppearsOnYourStatementAs string  `csv:"Appears On Your Statement As"`
	Address                  string  `csv:"Address"`

	CityState `csv:"City/State"`

	ZipCode   string `csv:"Zip Code"`
	Country   string `csv:"Country"`
	Reference string `csv:"Reference"`

	Category `csv:"Category"`
}

type RawTransaction struct {
	Fields
	UserField string
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse("01/02/2006", string(data))
	if err != nil {
		return err
	}

	d.Time = t
	return nil
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
