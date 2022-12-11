package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// representa todas as rotas
type Rota struct{
	Uri 					string
	Metodo 					string
	Funcao 					func(http.ResponseWriter, *http.Request)
	RequerAutentificacao 	bool
}


// coloca todas as rotas no Router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}