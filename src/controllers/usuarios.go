package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario usuario no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()

	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido %d", usuarioID)))
}

// GetUsuario pega usuario no banco
func GetUsuario(w http.ResponseWriter, r *http.Request) {

}

// GetUsuarios usuarios no banco
func GetUsuarios(w http.ResponseWriter, r *http.Request) {

}

// AtualizarUsuario atualiza usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

}

// DeletaUsuario usuario no banco
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {

}
