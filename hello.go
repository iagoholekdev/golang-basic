package main

import (
	"example/hello/helper"
	"example/hello/pessoa"
)

func main() {
	p := pessoa.Pessoa{}
	p.SetIdade(20).SetNome("Iago")
	helper.Imprimir(p)
}
