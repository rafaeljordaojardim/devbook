package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar gera router
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
