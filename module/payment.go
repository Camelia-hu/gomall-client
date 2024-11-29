package module

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Uid                       uint32  `json:"uid"`
	Money                     float32 `json:"money"`
	CreditCardNumber          string  `json:"creditCardNumber"`
	CreditCardCvv             int32   `json:"creditCardCvv"`
	CreditCardExpirationYear  int32   `json:"creditCardExpirationYear"`
	CreditCardExpirationMonth int32   `json:"creditCardExpirationMonth"`
}
