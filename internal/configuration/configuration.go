package configuration

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

var currentconfig Config

func InitConfig() {
	dsn := "host=localhost user=postgres password=Mindsrl dbname=animationtest port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	currentconfig = Config{DB: db}
}

func Current() Config {
	return currentconfig
}

func RouterInit() *mux.Router {
	router := mux.NewRouter()
	router.Use(requestCorsMiddleware())
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
