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
