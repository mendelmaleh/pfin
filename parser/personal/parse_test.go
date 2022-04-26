package personal

import (
	"testing"
	"time"

	"git.sr.ht/~mendelmaleh/pfin/parser/util"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	data := []byte(`
date    	amount	media	description
----    	------	-----	-----------
- some comment
2022-04-26	31.41	card	some purchase
`)

	txns := []Transaction{{Raw: Raw{
		Date:        util.Date{time.Date(2022, time.April, 26, 0, 0, 0, 0, time.UTC)},
		Amount:      31.41,
		Media:       "card",
		Description: "some purchase",
	}}}

	raw, err := Parse(data)
	if err != nil {
		t.Errorf("error parsing data: %s", err)
	}

	assert := assert.New(t)

	for i, v := range raw {
		assert.Equal(txns[i], v)
	}

}
