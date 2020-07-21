package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	saida := readFiles()
	fmt.Println(saida)
}

func readFiles() []string {
	var nomes []string

	arquivo, err := os.Open("nomes.txt")
	reader := bufio.NewReader(arquivo)
	for {
		linha, err := reader.ReadString('\n')

		linha = strings.TrimSpace(linha)
		if err == io.EOF {
			break
		}
		nomes = append(nomes, linha)
	}

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return nomes
	}

	return nomes
}
