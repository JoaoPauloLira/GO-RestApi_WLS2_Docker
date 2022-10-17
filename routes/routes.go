package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"
)

const (
	GET    string = "Get"
	POST          = "Post"
	PUT           = "Put"
	DELETE        = "Delete"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods(GET)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods(GET)

	r.HandleFunc("/api/personalidades", controllers.CriarUmaNovaPersonalidade).Methods(POST)

	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods(DELETE)

	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods(PUT)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
