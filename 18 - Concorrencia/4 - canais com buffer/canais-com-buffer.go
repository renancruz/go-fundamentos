package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Programando e GO!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
