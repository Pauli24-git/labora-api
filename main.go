package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	//Tarea.
	router.HandleFunc("/items", getItems).Methods("GET")
	//router.HandleFunc("/items/{id}", getItem)
	router.HandleFunc("/item", getItem).Methods("GET")

	//Más adelante.
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items", updateItem).Methods("PUT")
	router.HandleFunc("/items/", deleteItem).Methods("DELETE")

	//nuevo endpoint
	router.HandleFunc("/items/details", getDetails).Methods("GET")

	http.ListenAndServe(":8080", router)

}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Paula"},
	{ID: "2", Name: "Lucas"},
	{ID: "3", Name: "Misa"},
	{ID: "4", Name: "Rosario"},
	{ID: "5", Name: "Epik High"},
	{ID: "6", Name: "Paula"},
	{ID: "7", Name: "Misa"},
	{ID: "8", Name: "Rosario siempre estuvo cerca"},
	{ID: "9", Name: "Bokita"},
	{ID: "10", Name: "Burzaco"},
}

// func getItems(w http.ResponseWriter, r *http.Request) {

// 	json.NewEncoder(w).Encode(items)
// }

func getItems(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := request.URL.Query() // page - itemsPerPage

	// page := params["page"]
	// itemsPerPage := params["itemsPerPage"]

	page := params.Get("page")
	itemsPerPage := params.Get("itemsPerPage")

	if page == "0" {
		page = "1"
	}
	if itemsPerPage == "0" {
		itemsPerPage = "3"
	}

	pageIndex, _ := strconv.Atoi(page)
	itemsPerPageInt, _ := strconv.Atoi(itemsPerPage)

	var newListItems []Item

	init := itemsPerPageInt * (pageIndex - 1)
	limit := init + itemsPerPageInt

	nroPage := float64(len(items)) / float64(itemsPerPageInt)
	nroPage = math.Ceil(nroPage)

	if pageIndex <= int(nroPage) {
		if limit > len(items) {
			newListItems = items[init:]
		} else {
			newListItems = items[init:limit]
		}
	}
	// Función para obtener todos los elementos
	json.NewEncoder(response).Encode(newListItems)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	v := query.Get("id")
	n := query.Get("name")

	encontrado := false

	for _, item := range items {
		if (item.ID) == v || strings.ToLower(item.Name) == strings.ToLower(n) {
			encontrado = true
			json.NewEncoder(w).Encode(item)
		}
	}

	if !encontrado {
		json.NewEncoder(w).Encode("No se encontro ningun registro con ese id o nombre")
	}
	// params := mux.Vars(r)
	// for _, item := range items {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item.Name)
	// 	}
	// 	if item.ID == params["name"] {
	// 		json.NewEncoder(w).Encode(item.Name)
	// 	}
	// }
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var decodedItem Item
	err := json.NewDecoder(r.Body).Decode(&decodedItem)

	//no funciona, consultar despues porque no captura error o porque deja sumar cosas vacias
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic("panikeamo")
	}

	items = append(items, decodedItem)
	json.NewEncoder(w).Encode(items)

}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nombreActualizado Item
	err := json.NewDecoder(r.Body).Decode(&nombreActualizado)
	query := r.URL.Query()
	idBuscado := query.Get("id")

	encontrado := false

	if nombreActualizado.Name != "" {
		for _, item := range items {
			if (item.ID) == idBuscado {
				item.Name = nombreActualizado.Name
				encontrado = true
				json.NewEncoder(w).Encode(item)
			}
		}
	}

	if !encontrado {
		json.NewEncoder(w).Encode("No se encontro ningun registro con ese id o nombre")
	}

	if err != nil {
		json.NewEncoder(w).Encode("error")
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
}

func getDetails(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{}
	detailsChannel := make(chan ItemDetails, len(items))
	detailedItems := make([]ItemDetails, len(items))

	for _, item := range items {
		wg.Add(1) // Creamos el escucha, sin aun crearse la gorutina

		go func(id string) {
			defer wg.Done() //Completamos el trabajo del escucha, al final de esta ejecución
			detailsChannel <- getItemDetails(id)
		}(item.ID)
	}

	wg.Wait()
	close(detailsChannel)

	for details := range detailsChannel {
		detailedItems = append(detailedItems, details)
	}

	json.NewEncoder(w).Encode(detailedItems)

}

func getItemDetails(id string) ItemDetails {
	// Simula la obtención de detalles desde una fuente externa con un time.Sleep
	time.Sleep(100 * time.Millisecond)
	var foundItem Item
	for _, item := range items {
		if item.ID == id {
			foundItem = item
			break
		}
	}
	return ItemDetails{
		Item:    foundItem,
		Details: fmt.Sprintf("Detalles para el item %s", id),
	}
}

// func getItemDetails(id string) ItemDetails {
// 	// Simula la obtención de detalles desde una fuente externa con un time.Sleep
// 	time.Sleep(100 * time.Millisecond)
// 	var foundItem Item
// 	for _, item := range items {
// 		if item.ID == id {
// 			foundItem = item
// 			break
// 		}
// 	}
// 	return ItemDetails{
// 		Item:    foundItem,
// 		Details: fmt.Sprintf("Detalles para el item %s", id),
// 	}
// }
