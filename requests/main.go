package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	intro()

	command := readCmd()

	resolveCommand(command)
}

func intro() {
	fmt.Println("1- Monitorar Site")
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
		// o _ ( underline ) pode ser utilizado pra demostrar o desinteresse em uma das variáveis, assim não é necessário que eu trabalhe com a variável que eu não quero trabalhar
		resp, _ := monitor("http://github.com/vhsoares")
		fmt.Println(resp)
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

func monitor(site string) (*http.Response, error) {
	resp, err := http.Get(site)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
