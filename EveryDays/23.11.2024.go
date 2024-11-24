package EveryDays

import (
	"fmt"
	"math"
	"strings"
)

func Saturday_23_11_24_1() {

	var firstUint uint8 = 8
	var secondUint uint16 = 4065
	var thirdUint uint32 = 568979845
	var fourthUint uint64 = 18446744073709551614

	float := 123.879

	boolean := true

	var line string = "Это строка"

	fmt.Println("Переменная типа uint8, которая хранит от 0 до 255 -", firstUint)
	fmt.Println("Переменная типа uint16, которая хранит от 0 до 65535 -", secondUint)
	fmt.Println("Переменная типа uint32, которая хранит от 0 до 4294967295 -", thirdUint)
	fmt.Println("Переменная типа uint64, которая хранит от 0 до 18446744073709551614 -", fourthUint)
	fmt.Println("Переменная типа float, которая хранит вещественные числа -", float)
	fmt.Println("Переменная типа boolean, которая хранит логический тип -", boolean)
	fmt.Println("Переменная типа string, которая хранит строковый тип -", line)

}
func Saturday_23_11_24_2() {
	var choice string
	fmt.Println("Выберете фигуру (Треугольник, Круг, Квадрат)")
	fmt.Scanf("%s", &choice)
	switch strings.ToLower(choice) {
	case "треугольник":
		{
			triangleArea()
		}
	case "круг":
		{
			circleArea()
		}
	case "квадрат":
		{
			squreArea()
		}
	}
}

func triangleArea() {
	var (
		Basis  float64
		Height float64
	)
	fmt.Print("Введите основание тругольника :")
	fmt.Scan(&Basis)
	fmt.Println()
	fmt.Print("Введите выосту тругольника :")
	fmt.Scan(&Height)

	area := (Basis * Height) / 2

	fmt.Println(area)
}

func circleArea() {
	var radius float64
	fmt.Println("Ваедите радиус круга")
	fmt.Scan(&radius)
	Area := math.Pi * radius * radius
	fmt.Println("Area :", Area)
}

func squreArea() {
	var side int
	fmt.Println("Ваедите сторону квадрата")
	fmt.Scan(&side)
	Area := side * side
	fmt.Println("Area :", Area)
}

func Saturday_23_11_24_3() {
	fmt.Println("Конвертер")
	var CelsiusDegrees float64
	var Kilometers float64
	fmt.Print("Введите градусы цельсия : ")
	fmt.Scan(&CelsiusDegrees)
	fmt.Printf("\nОтвет в градусах по фаренгейту : ")
	ConvertDegrees(CelsiusDegrees)

	fmt.Println("Конвертер километров в мили")

	fmt.Print("Введите количество километров : ")
	fmt.Scan(&Kilometers)
	fmt.Println("Ответ в милях")
	convertMiles(Kilometers)

}

// num int - Число градусов цельсия
func ConvertDegrees(num float64) {
	var Fahrengheit float64
	Fahrengheit = (num * 9 / 5) + 32
	fmt.Printf("%0.2f", Fahrengheit)
	fmt.Println()
}

func convertMiles(kilometers float64) {
	var Miles float64
	Miles = kilometers / 1.609
	fmt.Printf("%0.2f", Miles)
	fmt.Println()
}
