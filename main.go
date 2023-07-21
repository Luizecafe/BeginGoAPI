//Pacote principal

package main

// Importação dos pacotes

import (
	"fmt"
	"net/http"
)

// Função principal

func main() {
	// Tratamento da requisição
	http.HandleFunc("/", helloHandler)
	// Inicialização do servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// Função que será chamada quando a rota "/" for acessada
// Ela recebe duas informações: o ResponseWriter e o Request.

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Escreve "Hello, World!" na resposta.
	fmt.Fprintf(w, "Hello, World!")
}
