package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/AFIF-ZILANI/simple-web-server/pkg/types"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		w.Write([]byte("Create new Student"))
	}
}
