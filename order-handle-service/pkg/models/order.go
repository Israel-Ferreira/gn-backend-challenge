package models

import (
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/pkg/dto"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName  string  `gorm:"column:customer_name"`
	CustomerEmail string  `gorm:"column:customer_email"`
	CustomerPhone string  `gorm:"column:customer_phone"`
	OrderAdress   Address `gorm:"foreignKey:OrderRefer"`
}

func ToOrder(orderDto dto.OrderDTO) *Order {
	return &Order{
		CustomerName:  orderDto.UserInfoAttr.Name,
		CustomerEmail: orderDto.UserInfoAttr.Email,
		CustomerPhone: orderDto.UserInfoAttr.Phone,
		OrderAdress: Address{
			City:         orderDto.AddressAttr.City,
			Street:       orderDto.AddressAttr.Street,
			Neighborhood: orderDto.AddressAttr.Neighborhood,
			ZipCode:      orderDto.AddressAttr.ZipCode,
			Uf:           orderDto.AddressAttr.Uf,
		},
	}
}
