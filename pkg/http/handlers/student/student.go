package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/AFIF-ZILANI/simple-web-server/pkg/storage"
	"github.com/AFIF-ZILANI/simple-web-server/pkg/types"
	"github.com/AFIF-ZILANI/simple-web-server/pkg/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// validate request body

		if err := validator.New().Struct(student); err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		lastId, err := storage.CreateStudent(student)

		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, err)
		}

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		response.WriteJSON(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}
