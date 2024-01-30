package util

import (
	"time"
)

// DateISO decodes ISO8601 dates
type DateISO struct {
	time.Time
}

func (d *DateISO) UnmarshalJSON(data []byte) error {
	return d.UnmarshalText(data[1 : len(data)-1])
}

func (d *DateISO) UnmarshalText(data []byte) error {
	t, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		return err
	}

	d.Time = t

	return nil
}

func (d *DateISO) String() string {
	return d.Time.Format("2006-01-02")
}

func (d DateISO) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func NewDateISO(s string) (d DateISO, err error) {
	err = d.UnmarshalText([]byte(s))
	return
}

type DateUS struct {
	time.Time
}

func (d *DateUS) UnmarshalJSON(data []byte) error {
	return d.UnmarshalText(data[1 : len(data)-1])
}

func (d *DateUS) UnmarshalText(data []byte) error {
	t, err := time.Parse("01/02/2006", string(data))
	if err != nil {
		return err
	}

	d.Time = t

	return nil
}
