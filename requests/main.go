package main

import (
	"fmt"
	"net/http"
	"os"

	// ao se adicionar um repositório do github é necessário dar um goget
	"github.com/fatih/color"
)

// ao iniciar uma array constante é necessário utilizar a sintaxe correta, sendo
// const ~nome da constante~ = [~tamanho do array~]~tipo do array~{~ valores do array~}
const sites = [2]string{"http://github.com/vhsoares", "http://github.com/rkarwinski"}

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
		for {
			monitor(sites)
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

func monitor(webpages []string) (*http.Response, error) {
	// para saber o tamanho de um array use o len
	for i :=0 < webpages.len {
		testAvailability(webpages[i])
		i++
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
