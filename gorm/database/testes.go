package database

import (
	"errors"

	"gorm/models"
)

func (d *Database) CreateTeste(strTeste string, entity *models.Teste) error {
	entityDb, _ := d.ReadTeste(strTeste)
	if entityDb != nil {
		return errors.New("entity already exists")
	}
	d.conn.Create(entity)
	return nil
}

func (d *Database) ReadTeste(strTeste string) (*models.Teste, error) {
	var entityDb models.Teste
	d.conn.First(&entityDb, "str_teste = ?", strTeste)
	if entityDb.StrTeste == "" {
		return nil, errors.New("entity not found")
	}
	return &entityDb, nil
}

func (d *Database) UpdateTeste(strTeste string, entity *models.Teste) error {
	entityDb, err := d.ReadTeste(strTeste)
	if entityDb == nil {
		return err
	}
	d.conn.Save(entity)
	return nil
}

func (d *Database) DeleteTeste(strTeste string) error {
	entityDb, _ := d.ReadTeste(strTeste)
	if entityDb != nil {
		d.conn.Delete(&entityDb, "str_teste = ?", strTeste)
	}
	return nil
}
