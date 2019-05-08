package middlewares

import (
	"log"
	"net/http"
)

// Cors wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := "X-Requested-With, Accept, Accept-Language, Content-Type, JWT_TOKEN, ONNI_MAGIC, X-AUTH-APPLICATION-TOKEN, X-AUTH-TOKEN, X-CLIENT-ID"
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT,PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", headers)
		}
		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			log.Println("OPTIONS")
			return
		}

		next.ServeHTTP(w, r)
	})
}
