package main

import (
	"log"

	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/bd"
	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/handlers"
)

func main() {
	if bd.CheckConnection() == false {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()

}
