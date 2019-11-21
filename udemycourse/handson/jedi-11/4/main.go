//Starting with this code(https://play.golang.org/p/wlEM1tgfQD), use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your
//lat "50.2289 N"
//long "99.4656 W"

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, &sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  fmt.Errorf("Cannot compute the square root of negative numbers: %v", f),
		}
	}
	return 42, nil
}
