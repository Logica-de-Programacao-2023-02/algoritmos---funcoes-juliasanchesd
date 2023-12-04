package main

import (
	"errors"
	"fmt"
)

func filterStringsByLength(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	var result []string

	for _, str := range slice {
		if len(str) > 5 {
			result = append(result, str)
		}
	}

	return result, nil
}

func main() {
	stringSlice := []string{"apple", "banana", "orange", "grape", "kiwi"}
	filteredStrings, err := filterStringsByLength(stringSlice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings com mais de 5 caracteres:", filteredStrings)
	}
}
