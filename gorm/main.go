package main

import (
	"fmt"

	"gorm/database"
	"gorm/models"
)

func main() {
	url := "host=localhost port=5432 user=postgres password=root dbname=teste sslmode=disable"
	db, _ := database.Open(url)
	db.Migrate()

	var strTeste = "654321789"
	db.DeleteTeste(strTeste)
	fmt.Printf("Delete: %v\n", strTeste)

	entity := &models.Teste{StrTeste: strTeste, IntTeste: 100, FloatTeste: 10.0}
	entity.SetBlobTeste([]string{"a", "b", "c"})
	db.CreateTeste(strTeste, entity)
	fmt.Printf("Create: %v\n", strTeste)

	entity.SetBlobTeste([]string{"x", "y", "z"})
	db.UpdateTeste(strTeste, entity)
	fmt.Println("Update: BlobTeste x, y, z")

	entity, err := db.ReadTeste(strTeste)
	fmt.Printf("Read: %v, %v\n", entity.StrTeste, err)
}
