package middlewares

import (
	"bytes"
	"context"
	"image"
	_ "image/jpeg" // import the image package in order to read image dimensions
	_ "image/png"
	"net/http"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
)

// GridFSRepoKey is the shit
var GridFSRepoKey key = "gridfs-repo"

// AttachGridFSCollection skjdfn
func AttachGridFSCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		fs, err := interfaces.NewGridFS(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		ctx := context.WithValue(r.Context(), GridFSRepoKey, fs)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

const mB = 1 << 20

// ImageConfigKey sdf
var ImageConfigKey key = "image-config"

// FileHeaderKey sdf
var FileHeaderKey key = "gridfs-file-header"

// ReadFileFromBody skjdfn
func ReadFileFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(2 * mB)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		file, _, err := r.FormFile("file")
		fileHeader := make([]byte, r.ContentLength)
		_, err = file.Read(fileHeader)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		config, _, _ := image.DecodeConfig(bytes.NewReader(fileHeader))
		fileCtx := context.WithValue(r.Context(), ImageConfigKey, config)
		ctx := context.WithValue(fileCtx, FileHeaderKey, fileHeader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
