package util

import (
	"bytes"
	"fmt"
	"strconv"
)

type Amount int

func (a *Amount) Cents() int {
	return int(*a)
}

func (a *Amount) Float64() float64 {
	return float64(*a) / 100
}

func (a *Amount) String() string {
	return fmt.Sprintf("%.2f", a.Float64())
}

func (a *Amount) UnmarshalText(data []byte) error {
	// NOTE: assumes that there are always two decimal places

	// remove digit separators
	data = bytes.Map(func(r rune) rune {
		if r == ',' || r == '.' {
			return -1
		}
		return r
	}, data)

	i, err := strconv.Atoi(string(data))
	*a = Amount(i)

	return err
}
