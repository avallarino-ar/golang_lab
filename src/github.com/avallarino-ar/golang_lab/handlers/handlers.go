package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/middlew"
	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers setea puerto y pone el servidor a escuchar */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
