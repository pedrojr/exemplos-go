package main

import "fmt"

type NumberTypes interface {
	int | int32 | int64 | float32 | float64
}

func somaLista[T NumberTypes](lista []T) T {
	var soma T
	for _, item := range lista {
		soma += item
	}
	return soma
}

func main() {
	fmt.Printf("Soma {1, 2, 3, 4, 5} = %d\n", somaLista([]int32{1, 2, 3, 4, 5}))
	fmt.Printf("Soma {1.1, 2.1, 3.1, 4.1, 5.1} = %.2f\n", somaLista([]float32{1.1, 2.1, 3.1, 4.1, 5.1}))
}
