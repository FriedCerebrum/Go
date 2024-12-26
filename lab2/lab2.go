package main

//Гойда

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	///////////////////////////////////////////////////////////////////////////
	// Задание 1.1: Функция formatIP

	fmt.Println("Задание 1.1: Функция formatIP")
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println("IP-адрес:", formatIP(ip))
	fmt.Println()

	///////////////////////////////////////////////////////////////////////////
	// Задание 1.2: Функция listEven

	fmt.Println("Задание 1.2: Функция listEven")
	evenNumbers, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Чётные числа в диапазоне [1, 10]:", evenNumbers)
	}
	fmt.Println()

	///////////////////////////////////////////////////////////////////////////
	// Задание 2: Подсчёт вхождений символов

	fmt.Println("Задание 2: Подсчёт вхождений символов")
	text := "Привет, мир!"
	counts := countChars(text)
	fmt.Println("Текст:", text)
	fmt.Println("Вхождения символов:")
	for char, count := range counts {
		fmt.Printf("Символ '%c' встречается %d раз(а)\n", char, count)
	}
	fmt.Println()

	///////////////////////////////////////////////////////////////////////////
	// Задание 3: Структуры, методы и интерфейсы

	fmt.Println("Задание 3: Структуры, методы и интерфейсы")

	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 3, Y: 4}
	segment := Segment{Start: p1, End: p2} // Сегмет це отрезок
	fmt.Printf("Длина отрезка: %.2f\n", segment.Length())
	fmt.Println()

	triangle := Triangle{
		A: Point{0, 0},
		B: Point{0, 4},
		C: Point{3, 0},
	}

	circle := Circle{
		Center: Point{0, 0},
		Radius: 5,
	}

	printArea(triangle)
	printArea(circle)
	fmt.Println()

}

///////////////////////////////////////////////////////////////////////////////
// Задание 1.1: Функция formatIP

func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

///////////////////////////////////////////////////////////////////////////////
// Задание 1.2: Функция listEven

func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница больше правой")
	}
	var evens []int //це срез
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

///////////////////////////////////////////////////////////////////////////////
// Задание 2: Подсчёт вхождений символов

func countChars(s string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	return charCount
}

///////////////////////////////////////////////////////////////////////////////
// Задание 3: Структуры, методы и интерфейсы

type Point struct {
	X float64
	Y float64
}

type Segment struct {
	Start Point
	End   Point
}

// Вычисление длины отрезка
func (s Segment) Length() float64 {
	dx := s.End.X - s.Start.X
	dy := s.End.Y - s.Start.Y
	return math.Sqrt(dx*dx + dy*dy) // Теорема Пифагора для вычисления длины отрезка
}

type Triangle struct {
	A Point
	B Point
	C Point
}

type Circle struct {
	Center Point
	Radius float64
}

// Площадь треугольника
func (t Triangle) Area() float64 {
	return math.Abs((t.A.X*(t.B.Y-t.C.Y) + t.B.X*(t.C.Y-t.A.Y) + t.C.X*(t.A.Y-t.B.Y)) / 2.0)
}

// Площадь круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Интерфейс "Shape"
type Shape interface {
	Area() float64
}

// Обращается к интерфейсу
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}
