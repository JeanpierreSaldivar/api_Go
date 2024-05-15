package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//Gerar va a retornanr un router con las rutas configuradas
func Gerar() *mux.Router {
	r:= mux.NewRouter()

	return rotas.Configurar(r)
}