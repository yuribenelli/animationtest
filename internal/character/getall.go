package character

import (
	"animationtest/internal/configuration"
	"animationtest/internal/model"
	"encoding/json"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var pgs []model.Characters
	db := configuration.Current().DB

	if err := db.Find(&pgs).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(&pgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
