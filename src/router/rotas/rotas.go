package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas las rutas de la API
type Rota struct {
	URI    				string
	Metodo 				string
	Funcao 				func(http.ResponseWriter, *http.Request)
	RequerAutenticacao 	bool
}

//Configurar coloca todas las rutas dentro del router
func Configurar(r *mux.Router) *mux.Router{
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}