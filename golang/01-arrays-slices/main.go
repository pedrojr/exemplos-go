package main

import (
	"fmt"
)

func main() {
	fmt.Println("*** Array int ***************")
	listInt := [6]int{1, 1, 2, 3, 5, 8}
	for i := 0; i < len(listInt); i++ {
		fmt.Println("Num: ", listInt[i])
	}
	fmt.Println("")

	fmt.Println("*** Slice string ************")
	semana := []string{"segunda", "terça", "quarta", "quinta", "sexta"}
	fmt.Println("Len: ", len(semana))
	fmt.Println("Cap: ", cap(semana))
	PrintSliceStr(semana)

	semana = append(semana, "sábado", "domingo")
	fmt.Println("Len: ", len(semana))
	fmt.Println("Cap: ", cap(semana))
}

func PrintSliceStr(array []string) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	fmt.Println("")
}
