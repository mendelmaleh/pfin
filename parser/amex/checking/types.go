package checking

import "git.sr.ht/~mendelmaleh/pfin/parser/util"

type Transaction struct {
	Fields // computed fields
	Raw    // json object
}

type Fields struct {
	User, Account string
}

type Raw struct {
	// can be fund transfer, check deposit or misc, description is generally more detailed
	Typename             string `json:"__typename"`
	Amount               `json:"amount"`
	Date                 util.DateISO `json:"date"`
	Description          string       `json:"description"`
	ExternalAccount      any          `json:"externalAccount,omitempty"`
	IsPartiallyAvailable any          `json:"isPartiallyAvailable,omitempty"`
	LifecycleStatus      string       `json:"lifecycleStatus"`
	Reference            string       `json:"reference"`
	RunningBalance       Amount       `json:"runningBalance"`
}

type Amount struct {
	Typename string  `json:"__typename"`
	Amount   float64 `json:"amount,string"`
	Currency string  `json:"currency"`
}

type Response struct {
	Data struct {
		ProductAccountByAccountNumberProxy struct {
			Typename          string `json:"__typename"`
			FinancialPosition struct {
				Typename     string `json:"__typename"`
				Transactions struct {
					Typename             string        `json:"__typename"`
					CheckingTransactions []Transaction `json:"checkingTransactions"`
				} `json:"transactions"`
			} `json:"financialPosition"`
		} `json:"productAccountByAccountNumberProxy"`
	} `json:"data"`
}
