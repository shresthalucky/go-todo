package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeRequestBody(r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&v)

	if err != nil {
		return err
	}

	return nil
}
