package character

import (
	"animationtest/internal/configuration"
	"animationtest/internal/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	var ch model.Characters
	//id paramaters from request
	rmap := mux.Vars(r)
	id := rmap["ID"]
	// db query
	db := configuration.Current().DB
	//connectionError
	if err := db.First(&ch, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// answering back
	if err := json.NewEncoder(w).Encode(&ch); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
