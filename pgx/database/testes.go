package database

import (
	"errors"
	"pgx/models"
)

func (d *Database) CreateTeste(strTeste string, entity *models.Teste) error {
	entityDb, _ := d.ReadTeste(strTeste)
	if entityDb != nil {
		return errors.New("entity already exists")
	}
	d.exec("insert into testes(str_teste, int_teste, float_teste, blob_teste) values($1, $2, $3, $4)",
		entity.StrTeste, entity.IntTeste, entity.FloatTeste, entity.BlobTeste)
	return nil
}

func (d *Database) ReadTeste(strTeste string) (*models.Teste, error) {
	var entityDb models.Teste
	d.getRecord(&entityDb, "select str_teste, int_teste, float_teste, blob_teste from testes where str_teste = $1", strTeste)
	if entityDb.StrTeste == "" {
		return nil, errors.New("entity not found")
	}
	return &entityDb, nil
}

func (d *Database) UpdateTeste(strTeste string, entity *models.Teste) error {
	return d.exec("update testes set int_teste = $1, float_teste = $2, blob_teste = $3 where str_teste = $4",
		entity.IntTeste, entity.FloatTeste, entity.BlobTeste, entity.StrTeste)
}

func (d *Database) DeleteTeste(strTeste string) error {
	return d.exec("delete from testes where str_teste = $1", strTeste)
}
