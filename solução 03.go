package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d"}
	r := cts(slice)
	fmt.Print("Seu slice concatenado fica ", r)
}

func cts(slice []string) string {
	soma := ""
	for _, l := range slice {
		soma += l
	}
	return soma
}
