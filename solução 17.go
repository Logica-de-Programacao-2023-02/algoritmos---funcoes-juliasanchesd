package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func concatenateAndSort(stringsSlice []string) (string, error) {
	if len(stringsSlice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(stringsSlice)

	result := strings.Join(stringsSlice, " ")

	return result, nil
}

func main() {
	stringsSlice := []string{"banana", "apple", "orange", "grape"}
	result, err := concatenateAndSort(stringsSlice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings em ordem alfabética:", result)
	}
}
