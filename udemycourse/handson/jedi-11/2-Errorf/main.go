//Start with this code (https://play.golang.org/p/9a1IAWy5E6). Create a custom error message using “fmt.Errorf”.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("Ocorreu um erro no parse do json: %v", err)
	}
	return bs, nil
}
