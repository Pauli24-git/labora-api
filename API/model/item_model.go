package model

type Item struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	Product    string `json:"product"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"totalPrice"`
}

var Vistas = 0

// func (item *Item) GetTotalPrice() int {
// 	return item.Price * item.Quantity
// }

// func (item *Item) GetTotalPrice() {
// 	item.TotalPrice = item.Price * item.Quantity
// }

func NewItem(id int, name string, date string, prod string, quantity int, price int) Item {
	precioTotal := quantity * price

	return Item{
		ID:         id,
		Name:       name,
		Date:       date,
		Product:    prod,
		Quantity:   quantity,
		Price:      price,
		TotalPrice: precioTotal,
	}
}
