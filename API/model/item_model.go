package model

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

var Vistas = 0
