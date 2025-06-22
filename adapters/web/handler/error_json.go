package handler

import "encoding/json"

func jsonError(msg string) []byte {
	errorResponse := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	r, err := json.Marshal(errorResponse)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
