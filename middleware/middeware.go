package middleware

import "net/http"

// Para que quando um dos endpoints seja chamados já sete antes no cabeçalho o retorno do json como tipo application/json
// Antes de chegar no controller a requisição passa por aq
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r) //recursividade, chamando todos os endpoints e setando o cabeçalho
	})
}
