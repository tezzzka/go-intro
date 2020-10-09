package main

import (
	"fmt"
	cars "github/tezzzka/go-intro/homework3"
)

func main() {
	Car := &cars.Car{
		AutoGeneral: cars.AutoGeneral{
			Vendor: "VW",
			Model:  "Touareg",
			Year:   2010,
		},
		Comment: "nice car",
	}
	// Пароль для метода `SetDriverName` Hello.
	Car.SetDriverName("Andrey", "Hello")

	fmt.Println(Car)

	//fmt.Println(Car.GetDriverName())

}
