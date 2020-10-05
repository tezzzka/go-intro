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

func Fibonacci(N int) {
	f0 := 0
	f1 := 1
	fmt.Print(f0, " ", f1, " ")
	for i := 2; i <= N; i++ {
		fmt.Print(f0+f1, " ")
		tmp := f0
		f0 = f1
		f1 = tmp + f1
	}
}
