package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	nvslice, err := anfuncao(slice, dobrar)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novo slice:", nvslice)
	}
}

func anfuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("O slice está vazio")
	}

	novoSlice := make([]int, len(slice))
	for i, valor := range slice {
		novoSlice[i] = funcao(valor)
	}
	return novoSlice, nil
}

func dobrar(x int) int {
	return x * 2
}
