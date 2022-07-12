package middleware

import "net/http"

// O middleware, neste caso, seta o content-type toda vez que um endpoint da aplicacao for chamado
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
