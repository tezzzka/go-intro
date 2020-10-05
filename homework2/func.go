package homework2

import (
	"errors"
	"fmt"
)

func EvenOdd(a int) (out bool, err error) {
	if a%2 == 0 {
		out = true
		return out, nil
	}
	err = errors.New("Нечет.")
	return false, err
}

func Del(a int) bool {
	if a%3 == 0 {
		return true
	}
	err := errors.New("Деление на 3 без остатка.")
	fmt.Println(err)
	return false
}

//Ф.Фибоначчи
func Fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

//Ф.Обертка для Фибоначчи-ф
func FibonacciWrapper(N int) {
	f := Fibonacci()
	//Например, добавим срез и будем туда пушить фибо-элементы. Например, чтобы передавать результат этой ф. в целом в другую ф.
	var intSlice []int
	for i := 0; i <= N; i++ {
		fmt.Print(f(), " ")
		intSlice = append(intSlice, f())
	}
	//Можно так сразу вывести срез
	//fmt.Println("***\n", intSlice)
}
