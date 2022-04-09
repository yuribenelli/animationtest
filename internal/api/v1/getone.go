package v1

import (
	"animationtest/internal/configuration"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getOne(element interface{}) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		//id paramaters from request
		rmap := mux.Vars(r)
		id := rmap["ID"]
		// db query
		db := configuration.Current().DB
		//connectionError
		if err := db.First(&element, id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// answering back
		if err := json.NewEncoder(w).Encode(&element); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
