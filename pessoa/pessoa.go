package pessoa

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) GetNome() string {
	return p.Nome
}

func (p *Pessoa) SetNome(nome string) *Pessoa {
	p.Nome = nome
	return p
}

func (p *Pessoa) GetIdade() int {
	return p.Idade
}

func (p *Pessoa) SetIdade(idade int) *Pessoa {
	p.Idade = idade
	return p
}
