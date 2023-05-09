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

func PostItem(item Item) (*int, error) {
	db, err := config.GetDatabase()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	var id int
	err = db.QueryRow("INSERT INTO items(customer_name, order_date, product, quantity, price) values ($1, $2, $3, $4, $5) RETURNING id", item.Name, item.Date, item.Product, item.Quantity, item.Price).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil

}

func UpdateItem(id int, nombre string) (bool, error) {
	db, err := config.GetDatabase()
	defer db.Close()

	if err != nil {
		return false, err
	}

	result, err := db.Exec("UPDATE items SET customer_name = $1 WHERE id =$2", nombre, id)

	resultado, err := result.RowsAffected()
	if resultado == 0 {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteItem(id int) (bool, error) {
	db, err := config.GetDatabase()
	defer db.Close()

	if err != nil {
		return false, err
	}

	result, err := db.Exec("DELETE FROM items WHERE id = $1", id)

	resultado, err := result.RowsAffected()
	if resultado == 0 {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}
