package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b", "c", "d"}
	r, err := ctscv(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Seu slice concatenado fica", r)
	}
}

func ctscv(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("Deu ruim")
	}

	soma := ""
	for _, l := range slice {
		soma += l + ","
	}
	return soma, nil
}
