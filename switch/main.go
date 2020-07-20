package main

import "fmt"

func main() {
	var command int

	fmt.Println("1- Fazer Requisição")
	fmt.Println("2- Mostrar Log")
	fmt.Println("3- Manter controle")

	fmt.Scan(&command)

	switch command {
	case 1:
		fmt.Println("Fazendo a Sua requisição")
	case 2:
		fmt.Println("Mostrando Log")
	case 3:
		fmt.Println("Saindo do sistema")
	default:
		fmt.Println("Comando inválido")
	}
}
