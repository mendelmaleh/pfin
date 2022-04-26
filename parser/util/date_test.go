package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateISO(t *testing.T) {
	d := DateISO{}

	if err := d.UnmarshalText([]byte("2022-04-26")); err != nil {
		t.Errorf("error unmarshaling text: %s", err)
	}

	assert := assert.New(t)
	assert.Equal(time.Date(2022, time.April, 26, 0, 0, 0, 0, time.UTC), d.Time)
}
