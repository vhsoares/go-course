package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	// ao se adicionar um repositório do github é necessário dar um goget
	"github.com/fatih/color"
)

// ao iniciar uma array variavel é necessário utilizar a sintaxe correta, sendo
// var ~nome da constante~ = [~tamanho do array~]~tipo do array~{~ valores do array~}
// Slices são abstrações de arrays, onde não é necessário dizer os indices e os limites
var sites = []string{"http://github.com/vhsoares", "http://github.com/rkarwinski"}

// slices são abstrações de arrays não sendo necessário definir o tamanho final
// em Go os arrays tem tipo
var sliceNomedeSites = []string{"github", "laravel"}

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

		// a função append adicona os valores no slice
		sites = append(sites, "http://github.com/leonardomacha")
		fmt.Println("Existem ", len(sites), " sendo testados. Podemos testar até ", cap(sites), " sites.")
		// loop infinito
		for {
			monitor(sites)
			fmt.Println("")
			time.Sleep(4 * time.Second)
		}
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

func monitor(webpages []string) {
	// foreach ou for range
	// o _ sempre pode ser usado quando não se tem interesse no valor de retorno de qualquer coisa.
	for _, value := range webpages {
		testAvailability(value)
	}

}

func testAvailability(site string) (*http.Response, error) {
	resp, err := http.Get(site)

	if resp.StatusCode == 200 {
		color.Green("Site %s carregado com sucesso!", site)
	} else {
		color.Red("Não foi possível acessar este site")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
