package main

import "fmt"

func main() {
	fmt.Println("1- Fazer Requisição")
	fmt.Println("2- Mostrar Log")
	fmt.Println("3- Manter controle")

	var command int
	fmt.Scan(&command)

	if command == 1 {
		fmt.Println("Sua requisição está sendo processada")
	} else if command == 2 {
		fmt.Println("Log Será exibido em breve")
	} else if command == 3 {
		fmt.Println("Não quero!!!")
	} else {
		fmt.Println("Comando Inválido, saíndo do sistema")
	}
}
