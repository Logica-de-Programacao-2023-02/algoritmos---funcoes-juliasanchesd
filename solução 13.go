package main

import (
	"errors"
	"fmt"
)

func sumOfDigits(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Número negativo não é permitido")
	}

	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum, nil
}

func main() {
	num := 12345
	sum, err := sumOfDigits(num)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos dígitos de", num, "é", sum)
	}
}
