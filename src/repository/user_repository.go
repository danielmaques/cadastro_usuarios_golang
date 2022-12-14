package repository

import (
	"database/sql"
	"fit_api/src/models"
)

type Usuarios struct {
	db *sql.DB
}

func NewRequestUsuarios(db *sql.DB) *Usuarios{
	return &Usuarios{db}
}

func (repository Usuarios) Criar(usuario models.UserModel) (uint64, error){
	statement, erro := repository.db.Prepare("insert into usuarios (nome, nick, email, senha) velues(?, ?, ?, ?)",)

	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	resultado , erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha,)
	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}