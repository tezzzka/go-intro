package main

import (
	"fmt"

	hm "github/tezzzka/go-intro/homework2"
)

func main() {

	X, e := hm.EvenOdd(4)
	Y := hm.Del(6)
	if e != nil {
		panic(e)
	}
	fmt.Println("Чет/нечет true/false: ", X)
	fmt.Println("Деление на 3 true/false: ", Y)

	hm.Fibonacci(100)

}
