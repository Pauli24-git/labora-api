package service

import (
	"errors"
	"labora-api/API/config"
	"labora-api/API/model"
	"strconv"
	"sync"
	"time"
)

var items = []model.Item{}

func ObtainItems() ([]model.Item, error) {
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

		// item := model.Item{ID: id, Name: name, Date: date.Format("2006-01-02"), Product: product, Quantity: quantity, Price: price}
		item := model.NewItem(id, name, date.Format("2006-01-02"), product, quantity, price)
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

func ObtainItem(idBuscado string, wg *sync.WaitGroup, m *sync.Mutex) (*model.Item, error) {
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
	var totalPrice int

	err = db.QueryRow("SELECT * FROM items WHERE id=$1", idBuscadoConvertido).Scan(&id, &name, &date, &product, &quantity, &price, &totalPrice)

	if err != nil {
		return nil, err
	}

	m.Lock()
	model.Vistas += 1
	defer m.Unlock()

	item := model.Item{ID: id, Name: name, Date: date.Format("2006-01-02"), Product: product, Quantity: quantity, Price: price, TotalPrice: totalPrice}
	wg.Done()
	return &item, err
}

func CreateNewItem(item model.Item) (*int, error) {
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
