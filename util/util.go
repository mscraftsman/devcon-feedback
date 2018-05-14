package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse represents json responses sent to client
type JSONResponse struct {
	Status bool        `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

// JSONOutputError outputs an error message with json content
func JSONOutputError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	msg, _ := json.Marshal(message)
	fmt.Fprint(w, `{"status": false, "error": `+string(msg)+`}`)
}

// JSONOutputResponse outputs a json response
func JSONOutputResponse(w http.ResponseWriter, data interface{}) {
	var response JSONResponse

	response.Status = true
	response.Data = data

	output, err := json.Marshal(response)
	if err != nil {
		JSONOutputError(w, http.StatusInternalServerError, "An error occurred.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}
