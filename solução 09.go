package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Esta é uma frase de exemplo"
	palavras, err := extrairPalavras(str)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras na string:", palavras)
	}
}

func extrairPalavras(str string) ([]string, error) {
	if str == "" {
		return nil, fmt.Errorf("A string está vazia")
	}

	palavras := strings.Fields(str)
	return palavras, nil
}
