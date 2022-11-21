package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/internal/services"
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/pkg/dto"
)

func OrderHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var orderDTO dto.OrderDTO

	if err := json.NewDecoder(r.Body).Decode(&orderDTO); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := services.SaveOrderInDatabase(orderDTO); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
