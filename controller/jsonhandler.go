package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MalformedRequest struct {
	status  int
	message string
}

const MAX_MB = 1

func (mr *MalformedRequest) Error() string {
	return mr.message
}
func (mr *MalformedRequest) Status() int {
	return mr.status
}

func DecodeJsonBody(w http.ResponseWriter, req *http.Request, dst interface{}) error {
	header := req.Header.Get("Content-Type")
	if header == "" || header != "application/json" {
		message := "Content-Type Header is not application/json"
		return &MalformedRequest{status: http.StatusUnsupportedMediaType, message: message}
	}

	req.Body = http.MaxBytesReader(w, req.Body, MAX_MB*1000000)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)

	if err != nil {
		var SyntaxError *json.SyntaxError
		var UnmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.Is(err, io.EOF):
			message := fmt.Sprintf("Request body must not be empty")
			return &MalformedRequest{status: http.StatusBadRequest, message: message}

		case strings.Contains(err.Error(), "request body too large"):
			message := fmt.Sprintf("Request body must not be larger than %dMB", MAX_MB)
			return &MalformedRequest{status: http.StatusRequestEntityTooLarge, message: message}

		case errors.As(err, &SyntaxError):
			message := fmt.Sprintf("Request body contans badly-formed JSON (at position) %d", SyntaxError.Offset)
			return &MalformedRequest{status: http.StatusBadRequest, message: message}

		case errors.Is(err, io.ErrUnexpectedEOF):
			message := fmt.Sprintf("Request body contains badly-formed JSON")
			return &MalformedRequest{status: http.StatusBadRequest, message: message}

		case errors.As(err, &UnmarshalTypeError):
			message := fmt.Sprintf("Request body contains invalid value for the %q field (at position %d)", UnmarshalTypeError.Field, UnmarshalTypeError.Offset)
			return &MalformedRequest{status: http.StatusBadRequest, message: message}

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			message := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{status: http.StatusBadRequest, message: message}

		default:
			return err
		}

	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		message := fmt.Sprintf("Request body must only contain a single JSON Object")
		return &MalformedRequest{status: http.StatusBadRequest, message: message}
	}

	return nil

}
