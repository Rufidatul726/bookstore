package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, data interface{}) error {
	if body, err := io.ReadAll(r.Body); err != nil {
		return err
	} else {
		if err := json.Unmarshal([]byte(body), data); err != nil {
			return err
		}
	}

	return nil
}
