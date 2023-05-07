package controller

import (
	"encoding/json"
	"labora-api/API/service"
	"net/http"
)

func GetItems(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	items, err := service.ObtainItems()

	if err != nil {
		json.NewEncoder(response).Encode("Hubo un error en la consulta:" + err.Error())
		return
	}

	// Función para obtener todos los elementos
	json.NewEncoder(response).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	v := query.Get("id")

	item, err := service.ObtainItem(v)
	if err != nil {
		json.NewEncoder(w).Encode("Hubo un error en la consulta " + err.Error())
		return
	}

	json.NewEncoder(w).Encode(item)
}

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	var decodedItem Item
// 	err := json.NewDecoder(r.Body).Decode(&decodedItem)

// 	//no funciona, consultar despues porque no captura error o porque deja sumar cosas vacias
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic("panikeamo")
// 	}

// 	items = append(items, decodedItem)
// 	json.NewEncoder(w).Encode(items)

// }

// func UpdateItem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var nombreActualizado Item
// 	err := json.NewDecoder(r.Body).Decode(&nombreActualizado)
// 	query := r.URL.Query()
// 	idBuscado := query.Get("id")

// 	encontrado := false

// 	if nombreActualizado.Name != "" {
// 		for _, item := range items {
// 			if (item.ID) == idBuscado {
// 				item.Name = nombreActualizado.Name
// 				encontrado = true
// 				json.NewEncoder(w).Encode(item)
// 			}
// 		}
// 	}

// 	if !encontrado {
// 		json.NewEncoder(w).Encode("No se encontro ningun registro con ese id o nombre")
// 	}

// 	if err != nil {
// 		json.NewEncoder(w).Encode("error")
// 	}
// }

// func DeleteItem(w http.ResponseWriter, r *http.Request) {
// 	// Función para eliminar un elemento
// }

// func GetDetails(w http.ResponseWriter, r *http.Request) {
// 	// Función para obtener el detalle de un item especifico
// }
