package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Chama os me metodos para validar e formatar os dados do usario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (Usuario *Usuario) validar(etapa string) error {
	if Usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if Usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if Usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(Usuario.Email); erro != nil {
		return errors.New("O email é inserido é invalido")
	}

	if etapa == "cadastro" && Usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)

	}

	return nil
}
