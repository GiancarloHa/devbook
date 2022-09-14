package controllers

import "net/http"

//CriarUsuario insere um usario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}
//BuscarUsuarios busca todos os usuarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}
//BuscarUsuario busca um usuario salvos no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuario!"))
}
//AtualizarUsuario atualiza um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario!"))
}
//DeletarUsuario atualiza um usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario!"))
}