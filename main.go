// первой задачи: Создай программу, которая проверяет список студентов и выводит, кто сдал экзамен.
// Тема: переменные, функции, if, for, мапы.
package main

import // импортируем два пакета – bufio и os
// с этим пакетом мы уже знакомы
(
	"fmt"
)

func main() {
	var x, y float32
	fmt.Println("Введите два числа!")
	fmt.Scan(&x, &y)

	var simbol string
	fmt.Println("Введите операцию для чисел (+, -, *, /, %):")
	fmt.Scan(&simbol)

	switch simbol {
	case "+":
		fmt.Printf("Результат: %.2f\n", x+y)
	case "-":
		fmt.Printf("Результат: %.2f\n", x-y)
	case "*":
		fmt.Printf("Результат: %.2f\n", x*y)
	case "/":
		if y == 0 || x == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Printf("Результат: %.2f\n", x/y)
		}
	case "%":
		fmt.Printf("Результат: %d\n", int(x)%int(y))
	default:
		fmt.Println("Ошибка: неизвестная операция")
	}
}

