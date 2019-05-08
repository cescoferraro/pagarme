package interfaces

import "net/http"

// ONNiMagic skjdfn
func ONNiMagic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("ONNI_MAGIC") != "kerpenmuitogostoso" {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
