package main

import (
	"errors"
	"fmt"
	"strconv"
)

func binaryProc(X uint16) uint16 {
	var output uint16
	for i := 0; i < 16; i++ {
		bit := (X >> i) & 1
		if i%2 == 1 {
			bit ^= 1
		}
		output |= (bit << i)
	}
	return output
}

func toBinary(x uint16) string {
	return fmt.Sprintf("%016s", strconv.FormatUint(uint64(x), 2))
}

func checkerDiapason(X uint16, err error) error {
	if err != nil || X > 65535 {
		return errors.New("Ошибка: число выходит за область определеннхы значений")
	}
	return nil
}

func main() {
	var X, output uint16
	fmt.Println("Введите число в диапазоне от 0 до 65535: ")
	_, err := fmt.Scan(&X)
	checkerDiapason(X, err)
	output = binaryProc(X)
	fmt.Printf("%d\n", output)
	fmt.Printf("%s\n", toBinary(output))
}
