package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Vitor"
	var idade = 25
	// não precisa de var
	versao := 1.2
	fmt.Println("String:", nome, " Idade :", idade, " Versão: ", versao)
	fmt.Println(reflect.TypeOf(nome))
	fmt.Println(reflect.TypeOf(idade))
	fmt.Println(reflect.TypeOf(versao))
}
