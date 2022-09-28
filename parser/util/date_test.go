package util

import (
	"testing"
	"time"
)

func TestDateISO(t *testing.T) {
	d := DateISO{}

	if err := d.UnmarshalText([]byte("2022-04-26")); err != nil {
		t.Errorf("error unmarshaling text: %s", err)
	}

	if c := time.Date(2022, time.April, 26, 0, 0, 0, 0, time.UTC); !c.Equal(d.Time) {
		t.Errorf("failed: expected %v, got %v", c, d.Time)
	}
}
