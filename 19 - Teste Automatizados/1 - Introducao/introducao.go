package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEnderaco("Rodovia dos Imigrantes")
	fmt.Println(TipoDeEndereco)
}
