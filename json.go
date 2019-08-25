package response

import (
	"encoding/json"
	"net/http"

	"github.com/dmitrymomot/go-errors"
)

// JSON http response
func JSON(w http.ResponseWriter, code int, data interface{}) error {
	var body M
	switch data.(type) {
	case errors.Error:
		body = M{"error": data}
	case M:
		body = data.(M)
	default:
		body = M{"data": data}
	}

	b, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, err = w.Write(b)

	return err
}

// JSONErr http response
func JSONErr(w http.ResponseWriter, err errors.Error) error {
	return JSON(w, err.GetCode(), err)
}
