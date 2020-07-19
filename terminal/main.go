package main

import "fmt"

func main() {
	nome := "Vítor"
	fmt.Println("Bem Vindo ", nome)
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	var comando2 int

	// o & identifica o endereço da variável
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi ", comando)

	fmt.Println("Escolha outro comando")

	// scan não precisa do modificador de decimal ou de string, ele vê qual o tipo de variavel ele espera
	fmt.Scan(&comando2)
	fmt.Println("O comando escolhido foi ", comando2)

}
