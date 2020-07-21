package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	saveFiles("Vitor")
}

func saveFiles(nome string) {
	arquivo, err := os.OpenFile("nomes.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err == os.ErrNotExist {
		fmt.Println(err)
	}
	// sobrescreve a linha
	arquivo.WriteString(time.Now().Format(time.RFC822) + nome + "\n")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.Close()
}
