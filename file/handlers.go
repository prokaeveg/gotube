package file

import (
	"fmt"
	"gotube/api"
	"net/http"
	"os"
)

const MaxUploadSize int64 = 1024 * 1024 * 1024

func HandleUploadedFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			api.RespondError(w, http.StatusInternalServerError, "Unable to parse form")
		}

		file, handler, err := r.FormFile("file")
		defer file.Close()

		if err != nil {
			api.RespondError(w, http.StatusInternalServerError, err.Error())
		}

		if handler.Size > MaxUploadSize {
			api.RespondError(w, http.StatusBadRequest, "File too large, max 1GB")
		}

		err = os.MkdirAll("./upload", os.ModePerm)
		if err != nil {
			api.RespondError(w, http.StatusInternalServerError, err.Error())
		}

		//dst, err := os.Create(fmt.Sprintf("./upload/%d%s", handler.Filename))

		fmt.Fprintf(w, "Uploaded File: %+v\n", handler.Filename)
		fmt.Fprintf(w, "File Size: %+v\n", handler.Size)
		fmt.Fprintf(w, "MIME Header: %+v\n", handler.Header.Get("Content-Type"))

	}
}
