package service

import (
	"labora-api/API/model"
)

func ObtainItems() ([]model.Item, error) {
	items, err := model.GetAllItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func ObtainItem(id string) (*model.Item, error) {
	item, err := model.GetSingleItem(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func CreateNewItem(item model.Item) (*int, error) {
	id, err := model.PostItem(item)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func UpdateItem(id int, nombre string) (bool, error) {
	encontrado, err := model.UpdateItem(id, nombre)
	if err != nil {
		return false, err
	}
	return encontrado, nil
}
