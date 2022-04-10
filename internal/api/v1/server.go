package v1

import (
	"animationtest/internal/model"
	"net/http"

	"github.com/gorilla/mux"
)

func RouterInit() *mux.Router {
	router := mux.NewRouter()
	router.Use(requestCorsMiddleware())
	router.HandleFunc("/character/{ID}", getOne(model.Characters{}))
	router.HandleFunc("/power/{ID}", getOne(model.Powers{}))
	router.HandleFunc("/character", getAll([]model.Characters{}))
	router.HandleFunc("/power", getAll([]model.Powers{}))
	return router

}

// RequestCorsMiddleware is a middleware enabling cors each request arrive to web server
func requestCorsMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, req)
		})
	}
}
