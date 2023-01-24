package main

import (
	"calculadora"
	"fmt"
)

func main() {
	fmt.Println("Soma: ", calculadora.Soma(15, 3))
	fmt.Println("Subtração: ", calculadora.Subtracao(15, 3))
	fmt.Println("Multiplicação: ", calculadora.Multiplicacao(15, 3))
	fmt.Println("Divisão: ", calculadora.Divisao(15, 3))
}
