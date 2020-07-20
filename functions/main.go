package main

import (
	"fmt"
	"os"
)

func main() {

	intro()

	command := readCmd()

	resolveCommand(command)
}

func intro() {
	fmt.Println("1- Fazer Requisição")
	fmt.Println("2- Mostrar Log")
	fmt.Println("3- Manter controle")
}

func readCmd() int {
	var command int
	fmt.Scan(&command)
	return command
}

func resolveCommand(command int) {
	switch command {
	case 1:
		fmt.Println("Fazendo a Sua requisição")
	case 2:
		fmt.Println("Mostrando Log")
	case 3:
		fmt.Println("Saindo do sistema de controle")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido")
		os.Exit(-1)
	}
}
