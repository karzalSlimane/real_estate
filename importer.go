package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

// ImportDataFromCSV lit un fichier CSV et insère les données dans la base de données
func ImportDataFromCSV(filename string) {
	// Ouvrir le fichier CSV
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Lire le contenu du fichier
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Parcourir chaque ligne (en ignorant l'en-tête)
	for i, record := range records {
		if i == 0 {
			continue // Sauter l'en-tête
		}

		// Convertir les colonnes en types appropriés
		price := record[1]
		datetime := record[2]
		nbRooms, _ := strconv.Atoi(record[3])
		nbBaths, _ := strconv.Atoi(record[4])
		surfaceArea, _ := strconv.ParseFloat(record[5], 64)
		link := record[6]
		cityName := record[7]
		equipements := strings.Split(record[8], ",") // Liste d'équipements

		// Vérifier si la ville existe déjà dans la base de données
		var city Ville
		DB.Where("name = ?", cityName).FirstOrCreate(&city, Ville{Name: cityName})

		// Créer une annonce
		ad := Annonce{
			Title:       record[0],
			Price:       price,
			Datetime:    datetime,
			NbRooms:     nbRooms,
			NbBaths:     nbBaths,
			SurfaceArea: surfaceArea,
			Link:        link,
			CityID:      city.ID,
		}
		DB.Create(&ad)

		// Gérer les équipements
		for _, equipementName := range equipements {
			equipementName = strings.TrimSpace(equipementName) // Supprimer les espaces inutiles

			// Vérifier si l'équipement existe déjà dans la base de données
			var equipement Equipement
			DB.Where("name = ?", equipementName).FirstOrCreate(&equipement, Equipement{Name: equipementName})

			// Créer la relation Annonce <-> Equipement
			annonceEquipement := AnnonceEquipement{
				AnnonceID:    ad.ID,
				EquipementID: equipement.ID,
			}
			DB.Create(&annonceEquipement)
		}
	}

	log.Println("Data imported successfully!")
}
