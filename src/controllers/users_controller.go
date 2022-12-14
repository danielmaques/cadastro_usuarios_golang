package controllers

import (
	"encoding/json"
	"fit_api/src/database"
	"fit_api/src/models"
	"fit_api/src/repository"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		log.Fatal(erro)
	} 

	var usuario models.UserModel
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil{
		log.Fatal(erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repository.NewRequestUsuarios(db)
	usuarioID, erro := repository.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido", usuarioID)))
};

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando todos os usuários..."))
};

func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando usuário..."))
};

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando usuário..."))
};

func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando usuário..."))
}