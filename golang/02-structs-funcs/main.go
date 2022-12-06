package main

type ItemVenda struct {
	Produto       string
	ValorUnitario float64
	Quantidade    int
}

func (item *ItemVenda) ValorTotal() float64 {
	return item.ValorUnitario * float64(item.Quantidade)
}
