package routes

import (
	"alura-go-restapi/controllers"
	"alura-go-restapi/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware) // Middleware atua como um interceptor, e faz uma ação para cada requisicao
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidade", controllers.BuscarPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidade/{id}", controllers.BuscarPersonalidadePorId).Methods("Get")
	r.HandleFunc("/api/personalidade", controllers.CriarPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidade/{id}", controllers.ExcluirPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidade/{id}", controllers.EditarPersonalidade).Methods("Put")

	// sobe server na porta 8000 e permite CORS vindo de qualquer dominio (representado por *). Posso informar
	// atraves de dominio ou ip quais hosts especificos quero que acessem minha API
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
