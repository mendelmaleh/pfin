package citi

import "time"

type Transaction struct {
	Fields // computed fields, namespaced so they don't conflict with the interface methods
	Raw    // raw fields from the csv
}

type Fields struct {
	Card    string
	User    string
	Account string
}

type Raw struct {
	Description string
	Amount      float64

	Purchased time.Time
	Posted    time.Time

	Cardmember string
	Method     string
	Category   string
	Rewards    int

	// Country    string

	Type string
}
