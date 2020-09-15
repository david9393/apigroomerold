package routers

import (
	"encoding/json"
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
)

func GroomerLista(w http.ResponseWriter, r *http.Request) {
	var t models.Groomer
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	results, err := bd.ListaGroomers(t)
	if err != nil {
		http.Error(w, "Error al consultar los groomer"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
