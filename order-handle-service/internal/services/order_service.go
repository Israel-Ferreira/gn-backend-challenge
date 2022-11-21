package services

import (
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/internal/config"
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/internal/publishers"
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/pkg/dto"
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/pkg/models"
)

func SaveOrderInDatabase(orderDTO dto.OrderDTO) error {
	order := models.ToOrder(orderDTO)

	if txn := config.DbConn.Save(order); txn.Error != nil {
		return txn.Error
	}

	publishers.PublishAddress(order.OrderAdress)

	return nil
}
