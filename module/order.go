package module

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Uid          uint32 `json:"uid"`
	UserCurrency string `json:"userCurrency"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	OrderItems   []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"orderID"`
	Cost      float32 `json:"cost"`
	ProductId uint32  `json:"productId"`
	Quantity  int32   `json:"quantity"`
	Order     Order   `gorm:"foreignKey:OrderID;references:ID"`
}

type OrderReq struct {
	UserId       uint32         `json:"userId"`
	UserCurrency string         `json:"userCurrency"`
	Address      string         `json:"address"`
	Email        string         `json:"email"`
	OrderItems   []OrderItemReq `json:"orderItems"`
}

type OrderItemReq struct {
	Cost      float32 `json:"cost"`
	ProductId uint32  `json:"productId"`
	Quantity  int32   `json:"quantity"`
}
