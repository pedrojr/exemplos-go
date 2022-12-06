package main

import (
	"testing"
)

func TestItemVenda(t *testing.T) {
	item := ItemVenda{Produto: "Cerveja", ValorUnitario: 4.75, Quantidade: 6}
	valorTotal := item.ValorTotal()
	valorEsperado := 28.5

	if valorTotal != valorEsperado {
		t.Errorf("valorTotal %.2f valorEsperado %.2f", valorTotal, valorEsperado)
	}
}
