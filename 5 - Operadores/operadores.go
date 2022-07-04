package main

import "fmt"

func main() {
	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)

	fmt.Println(soma2)

	//FIM DO ARITMÉTICOS

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//FIM DOS RELACIONAIS

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println("----------------")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero++     // numero = numero + 1
	numero += 15 // numero = numero + 15

	fmt.Println(numero)

	numero--     // numero = numero -1
	numero -= 20 // numero = numero -20

	numero *= 3 // numero = numero * 3
	numero /= 2 // numero = numero / 2
	numero %= 3 // numero = numero % 3

	fmt.Println(numero)

	//FIM DOS OPERADORES UNÁRIOS

}
