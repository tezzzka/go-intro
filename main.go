package main

import (
	cars "github.com/tezzzka/go-intro/homework3"
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
	Car.SetDriverName("Andrey", "passw")

	//fmt.Println(Car)

	//fmt.Println(Car.GetDriverName())

}
