package main

import (
	"fmt"
)

func main() {
	// 	i := 0

	// 	for i < 10 {
	// 		i++
	// 		fmt.Println("Incrementando i", i)
	// 		time.Sleep(time.Second)
	// 	}

	// 	for j := 0; j < 10; j++ {
	// 		fmt.Println("Incrementando j", j)
	// 		time.Sleep(time.Second)
	// 	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, palavra := range "PALAVRA" {
		fmt.Println(indice, string(palavra))
	}

	mapa := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}

	for indice, pessoa := range mapa {
		fmt.Println(indice, pessoa)
	}

}
