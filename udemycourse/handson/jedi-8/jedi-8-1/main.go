package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
}

func main() {

	users := []user{
		{"Jhonatas"},
		{"Rosendo"},
	}

	fmt.Println(users)

	json, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(json))
}
