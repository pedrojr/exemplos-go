package models

import (
	"bytes"
	"encoding/gob"

	"gorm.io/gorm"
)

type Teste struct {
	gorm.Model
	StrTeste   string
	IntTeste   int
	FloatTeste float32
	BlobTeste  []byte
}

func (r *Teste) SetBlobTeste(data any) {
	buffer := &bytes.Buffer{}
	gob.NewEncoder(buffer).Encode(data)
	r.BlobTeste = buffer.Bytes()
}
