package capitalone

import "time"

type RawTransaction struct {
	TransactionDate Date    `csv:"Transaction Date"`
	PostedDate      Date    `csv:"Posted Date"`
	CardNumber      string  `csv:"Card No."`
	Description     string  `csv:"Description"`
	Category        string  `csv:"Category"`
	Debit           float64 `csv:"Debit,omitempty"`
	Credit          float64 `csv:"Credit,omitempty"`
}

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
