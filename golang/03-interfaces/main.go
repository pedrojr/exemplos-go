package main

import "fmt"

type ContaBancaria interface {
	Depositar(valor float64)
	Sacar(valor float64)
	Saldo() float64
}

//---

type Conta struct {
	saldo float64
}

func (c *Conta) Depositar(valor float64) {
	c.saldo = c.saldo + valor
}

func (c *Conta) Sacar(valor float64) {
	c.saldo = c.saldo - valor
}

func (c *Conta) Saldo() float64 {
	return c.saldo
}

//---

type ContaPoupanca struct {
	Conta
}

func (c *ContaPoupanca) Render() float64 {
	c.saldo = c.saldo + (c.saldo * 0.05)
	return c.saldo
}

//---

func MovimentarConta(conta ContaBancaria) {
	fmt.Println("Depositar(150)")
	conta.Depositar(150)
	fmt.Println("Saldo:", conta.Saldo())
	fmt.Println("Sacar(50)")
	conta.Sacar(50)
	fmt.Println("Saldo:", conta.Saldo())
}

func main() {
	fmt.Println("\nContaCorrente")
	var contaCorrente Conta
	MovimentarConta(&contaCorrente)

	fmt.Println("\nContaPoupanca")
	var contaPoupanca ContaPoupanca
	MovimentarConta(&contaPoupanca)

	fmt.Println("Render()")
	contaPoupanca.Render()
	fmt.Println("Saldo:", contaPoupanca.Saldo())
	fmt.Println()
}
