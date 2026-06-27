package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {

	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {

	var errMsg []string

	for _, e := range errs {
		switch e.ActualTag() {
		case "required":
			errMsg = append(errMsg, e.Field()+" is required")
		default:
			errMsg = append(errMsg, e.Field()+" is invalid")
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsg, ", "),
	}
}
