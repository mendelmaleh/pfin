package util

import "time"

// Date implements the encoding.TextUnmarshaler interface to unmarshal from an ISO8601 date
type Date struct {
	time.Time
}

func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}
