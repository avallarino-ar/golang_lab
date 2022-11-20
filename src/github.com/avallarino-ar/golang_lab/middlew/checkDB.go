package middlew

import (
	"net/http"

	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/bd"
)

/*CheckDB es el middlw que permite conocer el estado de la DB*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc { // el middleware recibe algo y retorna lo mismo
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == false {
			http.Error(w, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)

	}
}
