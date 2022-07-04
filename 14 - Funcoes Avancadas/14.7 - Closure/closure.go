package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {
	texto := "Dentra da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}
