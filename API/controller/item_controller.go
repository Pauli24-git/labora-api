package controller

import (
	"encoding/json"
	"labora-api/API/model"
	"labora-api/API/service"
	"net/http"
	"strconv"
	"sync"
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
	var wg sync.WaitGroup
	var m sync.Mutex
	query := r.URL.Query()
	v := query.Get("id")
	var item *model.Item
	var err error

	wg.Add(1)
	go func() {
		item, err = service.ObtainItem(v, &wg, &m)
	}()
	wg.Wait()

	if err != nil {
		json.NewEncoder(w).Encode("Hubo un error en la consulta " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(&item)
	json.NewEncoder(w).Encode(model.Vistas)

}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var decodedItem model.Item
	err := json.NewDecoder(r.Body).Decode(&decodedItem)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	id, err := service.CreateNewItem(decodedItem)

	json.NewEncoder(w).Encode("El item fue creado correctamente. El numero de id es:")
	json.NewEncoder(w).Encode(id)

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	idBuscado := query.Get("id")
	nombre := query.Get("nombre")

	encontrado := false
	idBuscadoConvertido, err := strconv.Atoi(idBuscado)
	encontrado, err = service.UpdateItem(idBuscadoConvertido, nombre)

	if err != nil {
		json.NewEncoder(w).Encode("error")
		return
	}

	if !encontrado {
		json.NewEncoder(w).Encode("No se encontro ningun registro con ese id")
	} else {
		json.NewEncoder(w).Encode("Actualizado correctamente")
	}

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idBuscado := query.Get("id")

	encontrado := false
	idBuscadoConvertido, err := strconv.Atoi(idBuscado)
	encontrado, err = service.DeleteItem(idBuscadoConvertido)

	if err != nil {
		json.NewEncoder(w).Encode("error")
		return
	}

	if !encontrado {
		json.NewEncoder(w).Encode("No se encontro ningun registro con ese id")
	} else {
		json.NewEncoder(w).Encode("Eliminado correctamente")
	}
}

// func GetDetails(w http.ResponseWriter, r *http.Request) {
// 	// Función para obtener el detalle de un item especifico
// }
