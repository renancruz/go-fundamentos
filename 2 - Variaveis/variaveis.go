package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	//inventendo os valores das variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}