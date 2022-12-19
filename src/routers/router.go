package rotas

import (
	"fit_api/src/routers/rotas"

	"github.com/gorilla/mux"
)

// retornar rotas configuradas
func Gerar() *mux.Router{
	r := mux.NewRouter()
	return rotas.Configurar(r)
}