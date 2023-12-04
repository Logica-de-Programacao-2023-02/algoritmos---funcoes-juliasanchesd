package main

import (
	"errors"
	"fmt"
)

func isNumberInSlice(number int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, n := range slice {
		if n == number {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	// Teste da função
	slice := []int{1, 2, 3, 4, 5}
	numberToFind := 3

	isPresent, err := isNumberInSlice(numberToFind, slice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else if isPresent {
		fmt.Println(numberToFind, "está presente no slice.")
	} else {
		fmt.Println(numberToFind, "não está presente no slice.")
	}
}
