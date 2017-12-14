package decorator

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, value interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

func ErrorJson(w http.ResponseWriter, r *http.Request, err *ServerError) {
	switch err.Code {
	case BadRequest:
		w.WriteHeader(http.StatusBadRequest)

	case NotFound:
		w.WriteHeader(http.StatusNotFound)

	case Unauthorized:
		w.WriteHeader(http.StatusUnauthorized)

	case Forbidden:
		w.WriteHeader(http.StatusForbidden)

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	JsonResponse(w, map[string]error{"error": err})
}
