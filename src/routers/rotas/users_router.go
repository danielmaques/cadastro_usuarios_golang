package rotas

import (
	"fit_api/src/controllers"
	"net/http"
)

// rota de usuarios
var rotasUsuarios = []Rota{
	{
		Uri:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutentificacao: false,
	},
	{
		Uri:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutentificacao: false,
	},
	{
		Uri:    "/usuarios/{id}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutentificacao: false,
	},
	{
		Uri:    "/usuarios/{id}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutentificacao: false,
	},
	{
		Uri:    "/usuarios/{id}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutentificacao: false,
	},
}