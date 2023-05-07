package model

import (
	"errors"
	"labora-api/API/config"
	"strconv"
	"time"
)

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

var items = []Item{}

func GetAllItems() ([]Item, error) {
	db, err := config.GetDatabase()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM items")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var date time.Time
		var product string
		var quantity int
		var price int
		err := rows.Scan(&id, &name, &date, &product, &quantity, &price)
		if err != nil {
			return nil, err
		}

		item := Item{id, name, date.Format("2006-01-02"), product, quantity, price}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("Error en el servicio de base de datos:" + err.Error())
	}

	return items, err
}

func GetSingleItem(idBuscado string) (*Item, error) {
	db, err := config.GetDatabase()
	defer db.Close()

	idBuscadoConvertido, err := strconv.Atoi(idBuscado)
	if err != nil {
		return nil, err
	}

	var id int
	var name string
	var date time.Time
	var product string
	var quantity int
	var price int

	err = db.QueryRow("SELECT * FROM items WHERE id=$1", idBuscadoConvertido).Scan(&id, &name, &date, &product, &quantity, &price)

	if err != nil {
		return nil, err
	}

	item := Item{id, name, date.Format("2006-01-02"), product, quantity, price}
	return &item, err
}
