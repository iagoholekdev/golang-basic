package helper

import (
	"example/hello/pessoa"
	"fmt"
)

func Imprimir(p pessoa.Pessoa) {
	switch {
	case p.GetIdade() > 10:
		fmt.Println("Idade do " + p.GetNome() + " Maior que 10")
	case p.GetIdade() == 10:
		fmt.Println("Idade do " + p.GetNome() + " Igual a 10")
	default:
		fmt.Println("Idade do " + p.GetNome() + " Menor que 10")
	}
}
