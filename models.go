package main

type Annonce struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Price       string
	Datetime    string
	NbRooms     int
	NbBaths     int
	SurfaceArea float64
	Link        string
	CityID      uint
	City        Ville        `gorm:"foreignKey:CityID"`
	Equipements []Equipement `gorm:"many2many:annonce_equipements;"`
}

type Ville struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Equipement struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type AnnonceEquipement struct {
	AnnonceID    uint
	EquipementID uint
}
