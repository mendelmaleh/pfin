package util

import (
	"strings"
)

type StringFilter struct {
	String string
	Map    map[string]bool

	Default bool
}

// Filter returns true if s should be filtered out (skipped)
func (f *StringFilter) Filter(s string) bool {
	if f.String == "" {
		return false
	}

	// init
	if f.Map == nil {
		f.Map = make(map[string]bool)
		f.Default = true

		for _, v := range strings.Split(f.String, ",") {
			if strings.HasPrefix(v, "!") {
				f.Default = false
				f.Map[strings.TrimPrefix(v, "!")] = true
			} else {
				f.Map[v] = false
			}
		}
	}

	if v, ok := f.Map[s]; ok {
		return v
	}

	// string not in map
	return f.Default
}
