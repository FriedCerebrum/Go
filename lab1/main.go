package main

import (
	"errors"
	"fmt"
)

// Функция hello, которая принимает строку name и возвращает строку "Привет, name!"
func hello(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}

// Функция printEven, которая принимает диапазон (два целых числа) и выводит все четные числа в этом диапазоне
func printEven(a, b int64) error {
	// Проверка, что левая граница не больше правой
	if a > b {
		return errors.New("левая граница больше правой")
	}
	// Печать четных чисел в диапазоне [a, b]
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	return nil
}

// Функция apply, которая выполняет действия с двумя числами
func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// Проверка на деление на ноль
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	// Тестирование функции hello
	fmt.Println(hello("Абобий")) // Ожидается: Привет, Василий!

	// Тестирование функции printEven
	fmt.Println("Четные числа в диапазоне [1, 10]:")
	if err := printEven(1, 10); err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("Четные числа в диапазоне [10, 1]:")
	if err := printEven(10, 1); err != nil {
		fmt.Println("Ошибка:", err) // Ожидается ошибка: левая граница больше правой
	}

	// Тестирование функции apply
	result, err := apply(3, 5, "+")
	if err == nil {
		fmt.Printf("3 + 5 = %.2f\n", result) // Ожидается: 8.00
	} else {
		fmt.Println("Ошибка:", err)
	}

	result, err = apply(7, 10, "*")
	if err == nil {
		fmt.Printf("7 * 10 = %.2f\n", result) // Ожидается: 70.00
	} else {
		fmt.Println("Ошибка:", err)
	}

	result, err = apply(3, 5, "#")
	if err == nil {
		fmt.Printf("3 # 5 = %.2f\n", result)
	} else {
		fmt.Println("Ошибка:", err) // Ожидается ошибка: действие не поддерживается
	}

	result, err = apply(10, 0, "/")
	if err == nil {
		fmt.Printf("10 / 0 = %.2f\n", result)
	} else {
		fmt.Println("Ошибка:", err) // Ожидается ошибка: деление на ноль
	}
}
