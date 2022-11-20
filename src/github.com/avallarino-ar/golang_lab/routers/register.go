package routers

import (
	"encoding/json"
	"net/http"

	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/bd"
	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/models"
)

/*Register es la func para crear el registro del usuario en la BD*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // el Body es un Stream (solo lo puedo leer una vez)
	if err != nil {
		http.Error(w, "Error en los datos de entrada "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 0 {
		http.Error(w, "Debe ingresar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, exist, _ := bd.CheckExistUser(t.Email)
	if exist == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, satus, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "Error al insertar usuario "+err.Error(), 400)
		return
	}

	if satus == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
