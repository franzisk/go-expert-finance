package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody é a estrutura que vai retornar o status da aplicação
type HealthBody struct {
	Status string `json:"status"`
}

// HealthCheck é a função que recebe a requisição para checar o status da aplicação
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	encoder := json.NewEncoder(w)

	var health = HealthBody{
		Status: "alive",
	}

	// Para retornar um JSON precisa fazer um Encode com o pacote json
	if err := encoder.Encode(health); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
