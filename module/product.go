package module

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       float32  `json:"price"`
	Categories  []string `gorm:"type:text"`
}
