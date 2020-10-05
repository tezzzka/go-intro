package main

import (
	"fmt"

	hm "github.com/tezzzka/go-intro/homework2"
)

func main() {

	X, e := hm.EvenOdd(4)
	Y := hm.Del(6)
	if e != nil {
		panic(e)
	}
	fmt.Println("Чет/нечет true/false: ", X)
	fmt.Println("Деление на 3 true/false: ", Y)

	//hm.Fib(100)
	hm.FibonacciWrapper(20)

}
