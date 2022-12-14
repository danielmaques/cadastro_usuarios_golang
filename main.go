package main

import (
	"fit_api/src/config"
	rotas "fit_api/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	// fmt.Println(config.StringConexaoBanco)
	fmt.Print("rodando na porta :",config.Porta)

	r := rotas.Gerar()
	log.Fatal(http.ListenAndServe(":5000", r))
}