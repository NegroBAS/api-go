package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"../models"
)

// DisplayMessage devuelve mensajes al cliente
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("No se pudo convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
