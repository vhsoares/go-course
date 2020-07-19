package main

import "fmt"

func main() {
	var nome string = "Vítor"
	var idade int
	var versao float32 = 1.03
	fmt.Println("Olá senhor ", nome, "Sua idade é ", idade)
	fmt.Println("Versão ", versao)
}

// go não permite que você deixe variáveis inutilizadas no código
// float pode ser 32 bits ou 64 bits
