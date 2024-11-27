package main

import (
	"fmt"
	"strconv"
)

var bigNum = 1000

//Удаление комментариев

func symbol() {
	var sym int64
	fmt.Println("Введите число и узнаете символ")
	fmt.Scan(&sym)
	fmt.Println(string(sym))
}

func numbers() {
	var num string
	fmt.Println("Введите число:")
	fmt.Scan(&num)
	a, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Ошибка преобразования строки", err)
		return
	}
	fmt.Printf("Введено число: %T %d\n", a, a)
	b := float64(a)
	fmt.Printf("Вы ввели число: %T %F\n", b, b)
}

func comparison() {
	var a, b int
	fmt.Println("Введите 2 числа:")
	fmt.Scan(&a, &b)
	c := bool(a > b)
	fmt.Print(c, "/n")
}

func calculateSum(a, b int) int {
	return a * b
}

func division() {
	var a, b float64
	fmt.Println("Введите два числа для деления:")
	fmt.Scan(&a, &b)
	if b == 0 {
		fmt.Println("Ошибка: деление на ноль невозможно")
		return
	}
	result := a / b
	fmt.Printf("Результат деления: %.2f\n", result)
}

func greet(name string) {
	fmt.Println("Как тебя зовут?")
	fmt.Scanln(&name)
	fmt.Printf("Привет, %s!\n", name)
}

func calculateAverage(a, b, c, d float64) float64 {
	return (a + b + c + d) / 4
}

func main() {
	fmt.Println("Как тебя зовут?")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Привет, %s!\n", name)
	var age int
	fmt.Println("Сколько тебе лет?")
	fmt.Scanln(&age)
	fmt.Printf("Ничего себе, целых %d!\n", age)
	var x, y int
	fmt.Println("Введите два числа для умножения: ")
	fmt.Scan(&x, &y)
	result := calculateSum(x, y)
	fmt.Printf("Результат умножения: %d\n", result)
	fmt.Println(bigNum)

	average := calculateAverage(1, 2, 3, 4)
	fmt.Printf("Среднее значение переменной: %.2f\n", average)

	symbol()
	numbers()
	comparison()
	division()
	greet(name)
}
