package main

import (
	"errors"
	"fmt"
	"math"
)

type shape interface {
	area() (float64, error)
}

type triangle struct {
	a, b, c float64
}

func (tri triangle) area() (float64, error) {
	if tri.a <= 0 || tri.b <= 0 || tri.c <= 0 {
		return 0, errors.New("Ошибка: стороны треугольника должны быть положительными")
	}
	if tri.a+tri.b < tri.c || tri.c+tri.b < tri.a || tri.a+tri.c < tri.b {
		return 0, errors.New("Ошибка: не выполняется неравенство треугольника")
	}
	p := (tri.a + tri.b + tri.c) / 2
	return math.Sqrt(p * (p - tri.a) * (p - tri.b) * (p - tri.c)), nil
}

type sector struct {
	r, a float64
}

func (s sector) area() (float64, error) {
	if s.a < 0 || s.a > 2*math.Pi {
		return 0, errors.New("Ошибка: угол кругового сектора должен быть в диапазоне от 0 до 360 градусов")
	}
	if s.r <= 0 {
		return 0, errors.New("Ошибка: радиус должен быть положительным")
	}
	return 0.5 * s.r * s.r * s.a, nil
}

type trapezoid struct {
	a, b, h float64
}

func (t trapezoid) area() (float64, error) {
	if t.a <= 0 || t.b <= 0 {
		return 0, errors.New("Ошибка: основания трапеции должны быть положительными")
	}
	if t.h <= 0 {
		return 0, errors.New("Ошибка: высота трапеции должна быть положительной")
	}
	return (t.a + t.b) * t.h / 2, nil
}

func checkerErr(err1, err2, err3 error) error {
	if err1 != nil || err2 != nil || err3 != nil {
		return errors.New("Ошибка в чтении данных")
	}
	return nil
}

func main() {
	var (
		tri triangle
		t   trapezoid
		s   sector
	)
	fmt.Println("Введите длину сторон треугольника")
	_, err1 := fmt.Scan(&tri.a, &tri.b, &tri.c)
	fmt.Println("Введите длину оснований и высоты трапеции")
	_, err2 := fmt.Scan(&t.a, &t.b, &t.h)
	fmt.Println("Введите длину радиуса и угол в градусах")
	_, err3 := fmt.Scan(&s.r, &s.a)

	checkerErr(err1, err2, err3)

	fmt.Println(tri.area())
	fmt.Println(t.area())
	fmt.Println(s.area())
}
