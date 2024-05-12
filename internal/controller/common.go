package controller

import (
	"encoding/json"
	"net/http"
)

func parseParams(params interface{}, w http.ResponseWriter, r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func ResponseJSON(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(body)
}

func ResponseError(w http.ResponseWriter, body interface{}, status int) {
	content := ErrorResp{Error: body}
	ResponseJSON(w, content, status)
}

type ErrorResp struct {
	Error interface{} `json:"error"`
}
