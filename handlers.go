package main

import (
	"encoding/json"
	"net/http"
)

func adHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Récupération des annonces
		var ads []Annonce
		DB.Preload("City").Preload("Equipements").Find(&ads)
		json.NewEncoder(w).Encode(ads)
	case "POST":
		// Création d'une annonce
		var ad Annonce
		if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		DB.Create(&ad)
		w.WriteHeader(http.StatusCreated)
	}
}
