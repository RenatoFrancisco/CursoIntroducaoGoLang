package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando digitado foi:", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento...")

	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testanndo site", i, ":", site)
			testarSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testarSite(site string) {
	resp, _ := http.Get(site)

	statusCode := resp.StatusCode
	if statusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", statusCode)
	}
}
