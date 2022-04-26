package util

import "time"

// DateISO implements the encoding.TextUnmarshaler interface to unmarshal from an ISO8601 date
type DateISO struct {
	time.Time
}

func (d *DateISO) UnmarshalText(data []byte) error {
	t, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

type DateUS struct {
	time.Time
}

func (d *DateUS) UnmarshalText(data []byte) error {
	t, err := time.Parse("01/02/2006", string(data))
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}
