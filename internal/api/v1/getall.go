package v1

import (
	"animationtest/internal/configuration"
	"animationtest/internal/model"
	"encoding/json"
	"net/http"
)

// getAll return a func that excute a select on a table(depending of type of arg) on DB and return a struct of data
func getAll(elements interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var pgs []model.Characters
		db := configuration.Current().DB

		if err := db.Find(&pgs).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(&pgs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
