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

// func main() {
// 	var x int
// 	fmt.Scan(&x)

// 	switch x {
// 	case 1:
// 		fmt.Println("Очень плохо")
// 	case 2:
// 		fmt.Println("Плохо")
// 	case 3:
// 		fmt.Println("Удовлетворительно")
// 	case 4:
// 		fmt.Println("Хорошо")
// 	case 5:
// 		fmt.Println("Отлично")
// 	default:
// 		fmt.Println("Некорректная оценка")
// 		// ваш код тут
// 	}
// }

// func main() {
// 	var pass string
// 	fmt.Scan(&pass)
// 	pass2 := utf8.RuneCountInString(pass)
// 	if pass2 <= 8 {
// 		fmt.Println("Легкий")
// 	} else if pass2 <= 12 {
// 		fmt.Println("Средний")
// 	} else if pass2 <= 16 {
// 		fmt.Println("Сложный")
// 	} else if pass2 < 16 {
// 		fmt.Println("Очень сложный")
// 	}
// 	// var a int
// 	// fmt.Scan(&a)

// 	// result := a >= 13 && a <= 19
// 	// fmt.Println(result)

// 	//fmt.Println(false || false || true || false || 10 < 9 && false)
// 	// var a, b, c int
// 	// fmt.Scan(&a, &b, &c)

// 	// // run := rune(a + b)
// 	// fmt.Printf("%c%c%c\n", rune(a), rune(b), rune(c))
// }

// func main() {
// 	// var v float64
// 	// var t int
// 	// fmt.Scan(&v, &t)
// 	// s := v * float64(t)
// 	// fmt.Printf("%.2f", s)
// 	var r rune = 'F' + 5
// 	fmt.Println(string(r))
// }

// func main() {
// 	var name string
// 	fmt.Scan(&name)
// 	fmt.Println("Hello,", name)
// }

// func main() {
// 	var a, b float64
// 	fmt.Scan(&a, &b)
// 	fmt.Printf("Сумма: %.3f\n", a+b)
// 	fmt.Printf("ВЫчитание: %.3f\n", a-b)
// 	fmt.Printf("Умножение: %.3f\n", a*b)
// 	fmt.Printf("Деление: %.3f\n", a/b)

// }

// func main() {
// 	var distance float64
// 	var time int
// 	fmt.Scan(&distance, &time)
// 	speed := distance / float64(time)
// 	fmt.Printf("Speed: %.2f\n", speed)
// // 	"%.2f\n"
// // Это форматная строка, в которой задается, как вывести значение, передаваемое в функцию. Вот что значат её части:

// // % — специальный символ, который указывает на место, где будет выведено значение (например, число или строка). Всё, что после %, указывает на тип данных и формат отображения.

// // .2f — это означает, что мы хотим вывести число с плавающей точкой (тип float64 или float32) с точностью до 2 знаков после запятой:

// // f — обозначает число с плавающей точкой (float).
// // .2 — указывает на количество знаков после запятой. В данном случае это 2 знака.
// // \n — это символ новой строки. Он указывает, что после вывода значения курсор перейдёт на новую строку. Это важно, чтобы вывод был аккуратным и не продолжался в той же строке.
// }

// func main() {
// 	var a, b float32
// 	fmt.Scan(&a, &b)
// 	fmt.Println(a + b)
// 	fmt.Println(a - b)
// 	fmt.Println(a * b)
// 	fmt.Println(a / b)
// }

// func main() {
// 	var a, b, c float64
// 	fmt.Scan(&a, &b, &c)

// 	sum := 22*a - 234*b*c*(5+12*c)/259*a*c
// 	fmt.Println(int(sum))

// }

// func main() {
// 	var m float64
// 	var s int
// 	fmt.Scan(&m, &s)
// 	speed := m / float64(s)
// 	fmt.Println(speed)
// }

// func main() {
// 	// a := 10
// 	b := 100 % 90 //(1 + 2 + 3 - 2 - 2) * 10 * (100 % 90)
// 	// c := a * 10
// 	// d := c * a / 2
// 	//e := d % b
// 	fmt.Println(b)
// }

// func main() {
// 	var a int = 1000
// 	var1 := int8(a)
// 	fmt.Println(var1)
// }

// func main() {
// 	var a string = "a"
// 	var b string
// 	b = "b"
// 	c := "c"
// 	var d string = a
// 	e := c
// 	var f = b
// 	fmt.Println(a, b, c, d, e, f)
// }

// func main() {
// 	var a string
// 	var b string = "milk"
// 	var c, d int
// 	var s, e int = 40, 89
// 	g := "true"
// 	h, t := true, "min"
// 	const j = "false"

// 	fmt.Println(a, b, c, d, s, e, g, h, t, j)
// 	// your code here
// }
