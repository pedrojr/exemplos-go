package models

import (
	"bytes"
	"encoding/gob"
)

type Teste struct {
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
