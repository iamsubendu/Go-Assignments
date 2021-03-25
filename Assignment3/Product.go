//Rest api with Mux
//Building an inventory system

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json: "UID"`
	Name  string  `json:"Name"`
	Desc  string  `json: "Desc"`
	Price float64 `json: "Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/ => worked i.e, homePage()")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "/inventory => getInventory() is called")
	fmt.Println("getInventory() worked")
	//here you can understand difference between Fprintf & Println

	json.NewEncoder(w).Encode(inventory)
	//as this gets encoded it will popup in our browser
}

func createInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)
	//appending item to slice inventory

	json.NewEncoder(w).Encode(item)
	//Marshal => String
	//Encode => Stream
	//Marshal and Unmarshal convert a string into JSON
	//and vice versa. Encoding and decoding convert a stream
	// into JSON and vice versa

}

func deleteInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUid(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			//delete item from slice
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

func updateInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	//delete the item
	_deleteItemAtUid(params["uid"])
	//create a new data
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(inventory)

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createInventory).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteInventory).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateInventory).Methods("PUT")

	log.Fatal(http.ListenAndServe(":2000", router))
}

func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "iuhsih",
		Desc:  "ugduhgui uyhgdugid igdiuhiuhdiohoisdiodh",
		Price: 5.897,
	})
	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "iugfghsih",
		Desc:  "ugduhgui dsds dsdsd uyhgdugid igdiuhiuhdiohoisdiodh",
		Price: 3.907,
	})

	handleRequests()
}
