package models

type AddressMsg struct {
	OrderID      uint64 `json:"order_id"`
	City         string `json:"city"`
	Street       string `json:"street"`
	ZipCode      string `json:"zipcode"`
	Uf           string `json:"uf"`
	Neighborhood string `json:"neighborhood"`
}

func ToAddressMsg(address Address) *AddressMsg {
	return &AddressMsg{
		OrderID:      uint64(address.OrderRefer),
		City:         address.City,
		Street:       address.Street,
		ZipCode:      address.ZipCode,
		Uf:           address.Uf,
		Neighborhood: address.Neighborhood,
	}
}
