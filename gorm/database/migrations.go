package database

import (
	"gorm/models"
)

func (d *Database) Migrate() {
	d.conn.AutoMigrate(&models.Teste{})
}
