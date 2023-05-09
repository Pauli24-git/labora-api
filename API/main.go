package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"labora-api/API/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", controller.GetItems).Methods("GET")
	router.HandleFunc("/item", controller.GetItem).Methods("GET")
	router.HandleFunc("/items", controller.CreateItem).Methods("POST")
	router.HandleFunc("/items", controller.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/", controller.DeleteItem).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

// 	// // page := params["page"]
// 	// // itemsPerPage := params["itemsPerPage"]

// 	// page := params.Get("page")
// 	// itemsPerPage := params.Get("itemsPerPage")

// 	// if page == "0" {
// 	// 	page = "1"
// 	// }
// 	// if itemsPerPage == "0" {
// 	// 	itemsPerPage = "3"
// 	// }

// 	// pageIndex, _ := strconv.Atoi(page)
// 	// itemsPerPageInt, _ := strconv.Atoi(itemsPerPage)

// 	// var newListItems []Item

// 	// init := itemsPerPageInt * (pageIndex - 1)
// 	// limit := init + itemsPerPageInt

// 	// nroPage := float64(len(items)) / float64(itemsPerPageInt)
// 	// nroPage = math.Ceil(nroPage)

// 	// if pageIndex <= int(nroPage) {
// 	// 	if limit > len(items) {
// 	// 		newListItems = items[init:]
// 	// 	} else {
// 	// 		newListItems = items[init:limit]
// 	// 	}
// 	// }
// 	// Función para obtener todos los elementos
// 	json.NewEncoder(response).Encode(items)
// }

// func deleteItem(w http.ResponseWriter, r *http.Request) {
// 	// Función para eliminar un elemento
// }

// type ItemDetails struct {
// 	Item
// 	Details string `json:"details"`
// }

// func getDetails(w http.ResponseWriter, r *http.Request) {
// 	wg := sync.WaitGroup{}
// 	detailsChannel := make(chan ItemDetails, len(items))
// 	detailedItems := []ItemDetails{}
// 	var errors []error

// 	for _, item := range items {
// 		wg.Add(1) // Creamos el escucha, sin aun crearse la gorutina

// 		go func(id string) {
// 			defer wg.Done() //Completamos el trabajo del escucha, al final de esta ejecución
// 			details, err := getItemDetails(id)
// 			if err == nil {
// 				detailsChannel <- details
// 			} else {
// 				errors = append(errors, err)
// 			}
// 		}(item.ID)
// 	}

// 	wg.Wait()
// 	close(detailsChannel)

// 	for details := range detailsChannel {
// 		detailedItems = append(detailedItems, details)
// 	}

// 	json.NewEncoder(w).Encode(detailedItems)
// 	json.NewEncoder(w).Encode(errors)

// 	//en este segundo encoder quise mostrar el slice de errores, pero por alguna razon no me deja, me los trae vacios
// 	//Por que razon podria ser? No le pude encontrar la vuelta
// }

// func getItemDetails(id string) (ItemDetails, error) {
// 	// Simula la obtención de detalles desde una fuente externa con un time.Sleep
// 	time.Sleep(100 * time.Millisecond)
// 	var foundItem Item
// 	var errorEncontrado error
// 	var ItemEncontrado bool

// 	for _, item := range items {
// 		idNumero, _ := strconv.Atoi(item.ID)
// 		if item.ID == id && idNumero%2 == 0 {
// 			ItemEncontrado = true
// 			foundItem = item
// 			break
// 		}
// 	}

// 	if !ItemEncontrado {
// 		errorEncontrado = errors.New("No se pudo obtener el item" + id)
// 	}

// 	return ItemDetails{
// 		Item:    foundItem,
// 		Details: fmt.Sprintf("Detalles para el item %s", id),
// 	}, errorEncontrado
// }
