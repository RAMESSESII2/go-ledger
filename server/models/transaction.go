package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	DebitAmt  int64  `json:"debitamt"`
	CreditAmt int64  `json:"creditamt"`
}
