package main

import (
	"encoding/json"
	"fmt"
)

type Preson struct {
	Name   string   `json:"name"`
	Phone  int      `json:"phone"`
	Age    int      `json:"age"`
	Gender string   `json:"gender"`
	City   []string `json:"city"`
}

func main() {

	user := &Preson{Name: "Brock", Phone: 9888888882, Age: 99, Gender: "Male", City: []string{"kondagaon", "bilaspur"}}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))

	user2 := `{"Name":"Bullu","Phone":9089990892,"Age":99,"Gender":"Male","City":["ambikapur","bisrampur"]}`
	data2 := &Preson{}

	json.Unmarshal([]byte(user2), data2)
	fmt.Println(data2)
}