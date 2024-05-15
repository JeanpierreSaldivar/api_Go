package modelos

import (
	"errors"
	"strings"
	"time"
)

//Usuario representa un usuario utilizando la red social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func(usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa) ; erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error{
	if usuario.Nome == "" {
		return errors.New("El nombre es obligatorio y no puede estar en blanco")
	}
	if usuario.Nick == "" {
		return errors.New("El nick es obligatorio y no puede estar en blanco")
	}
	if usuario.Email == "" {
		return errors.New("El email es obligatorio y no puede estar en blanco")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("El contraseña es obligatorio y no puede estar en blanco")
	}
	return nil
}

func (usuario *Usuario) formatar(){
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}