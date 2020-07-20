package main

import "fmt"

func main() {
	nomes := []string{"Vítor", "Leonardo", "Vinicius", "Dani", "RK"}

	// existem 2 tipos de iteradores
	// Eles podem iterar tanto dentro de arrays como dentro de maps

	// for com len para o tamanho, ou for normal definindo um valor final
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

	// for com range (o equivalente a um foreach do Javascript)
	for iterator, value := range nomes {
		fmt.Println(iterator)
		fmt.Println(value)
	}
}
