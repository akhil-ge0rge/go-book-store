package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(x)
}
