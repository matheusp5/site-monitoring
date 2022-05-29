package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for { // Para o codigo rodar infinitamente ate o usuario clicar em "Sair do programa"
		exibeIntroducao()
		comand := getComand()
		verifyComand(comand)
	}
}

func exibeIntroducao() {
	versao := 1.2
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("0- Sair do programa")
}

func getComand() int { // Depois do parenteses colocamos o que ele vai retornar
	var comand int
	fmt.Scan(&comand)
	return comand
}

func verifyComand(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Nao reconheço esse comando!")
	}
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := []string{
		"https://alura.com.br",               // Site teste
		"https://stackoverflow.com",          // Site teste
		"https://google.com",                 // Site teste
		"https://google.com/VaiDarErro",      // Site teste
		"https://google.com/VaiDaErroTambem", // Site teste
	}

	for i := 0; i <= len(sites)-1; i++ {
		resp, _ := http.Get(sites[i])
		// fmt.Println(resp);
		if resp.StatusCode == 200 { // Pega o Codigo de status da resposta http
			fmt.Println("Site:", sites[i], "foi carregado com sucesso")
		} else if resp.StatusCode != 200 {
			fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode);
		}
	}
	fmt.Println("");
}
