package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Println("Введите коэффициенты квадратного уравнения\n")
	fmt.Println("Введите а\n")
	fmt.Scan(&a)
	fmt.Println("Введите b\n")
	fmt.Scan(&b)
	fmt.Println("Введите c\n")
	fmt.Scan(&c)
	D := b*b - 4*(a*c)
	if D > 0 {
		var x1 float64
		var x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("Корни уравнения\n")
		fmt.Println("x1:" + fmt.Sprint(x1))
		fmt.Println("x2:" + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64
		x = -b / (2 * a)
		fmt.Println("Корень уравнения\n")
		fmt.Println("x:" + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Уравнение не имеет корней\n")
	}
}
