package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Test  string      `json:"test,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	res := Response{Data: data}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(res)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, error error) {
	res := Response{Error: error.Error(), Data: nil}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// json.NewEncoder(w).Encode(res)
	j, _ := json.Marshal(res)
	w.Write(j)
}
