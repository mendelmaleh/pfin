package bofa

import "time"

type Transaction struct {
	Fields // computed fields
	Raw    // csv fields
}

type Fields struct {
	User string
}

type Raw struct {
	Date           Date    `csv:"Date"`
	Description    string  `csv:"Description"`
	Amount         float64 `csv:"Amount,omitempty"`
	RunningBalance float64 `csv:"Running Bal."`
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
