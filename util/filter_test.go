package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringFilter(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		filter string
		test   string
		result bool
	}{
		{"", "mendel", false},
		{"mendel", "mendel", false},
		{"mendel", "levi", true},
	}

	for i, test := range cases {
		f := StringFilter{String: test.filter}
		assert.Equalf(test.result, f.Filter(test.test), "failed test[%d]: %v", i, cases[i])
	}
}
