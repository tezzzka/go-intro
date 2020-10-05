package homework1

import (
	"fmt"
	"math"
)

//Задание №1.
func Calc() float32 {

	const tail float32 = 82.32
	var rub int

	fmt.Print("Введите сумму в рублях: ")
	fmt.Scanln(&rub)
	fmt.Printf("Результат: $%.2f USD\n ", float32(rub)/tail)

	return 0

}

//Задание №2.
func Ma() float32 {
	var kat1 float64 = 5.00
	var kat2 float64 = 3.00

	fmt.Println("Математика. Задача №2")
	fmt.Printf("Площадь прям.треугольника = %.2f\n", (kat1*kat2)/2)
	fmt.Printf("Периметр прям.треугольника = %.2f\n ", (kat1 + kat2))
	fmt.Printf("Гипотенуза прям.треугольника = %.2f\n ", math.Sqrt(math.Pow(kat1, 2)+math.Pow(kat2, 2)))

	return 0
}
