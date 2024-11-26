package module

import "gorm.io/gorm"

type Usr struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
