package main

import (
	"errors"
	"fmt"
)

func POW(x float64, y int) float64 {
	if y == 0 {
		return 1
	}
	var current float64 = 1
	for i := 0; i < y; i++ {
		current = current * x
	}
	return current
}

func Abs(n float64) float64 {
	if n < 0 {
		return n * (-1)
	} else {
		return n
	}
}

func fact(n int) float64 {
	var current int = 1
	for i := 1; i <= int(n); i++ {
		current *= i
	}
	return float64(current)
}

func sum(x float64, E float64) float64 {
	var (
		SUM float64 = 0
		n   int     = 0
	)
	for Abs((((POW(-1, n))*fact(2*n))/((1-float64(2*n))*POW(fact(n), 2)*POW(4, n)))*POW(x, n)) > E {
		SUM += (((POW(-1, n)) * fact(2*n)) / ((1 - float64(2*n)) * POW(fact(n), 2) * POW(4, n))) * POW(x, n)
		n++
	}
	return SUM
}
func checkerErr(err error) error {
	if err != nil {
		return errors.New("Ошибка при вводе данных")
	}
	return nil
}

func main() {
	var x, E float64
	_, err := fmt.Scan(&x, &E)
	checkerErr(err)
	fmt.Print(sum(x, E))
}
