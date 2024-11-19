package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialisation de la base de données

	initDB()

	// Migration des modèles

	DB.AutoMigrate(&Annonce{}, &Ville{}, &Equipement{}, &AnnonceEquipement{})

	//import data from csv

	log.Println("Importing data from CSV...")
	ImportDataFromCSV("appartemetn.csv")

	// Configuration des routes
	http.HandleFunc("/ads", adHandler)

	// Démarrage du serveur
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
