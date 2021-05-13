package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

var (
	logger, _     = zap.NewProduction(zap.Fields(zap.String("type", "api handler")))
	ErrBadRequest = errors.New("Bad Request")
)

func render(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	switch v := body.(type) {
	case string:
		_ = json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: v,
		})
	case error:
		_ = json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: v.Error(),
		})
	case nil:
		// do nothing
	default:
		_ = json.NewEncoder(w).Encode(body)
	}
}
