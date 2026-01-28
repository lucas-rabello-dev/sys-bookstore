package model

type Livro struct {
	Titulo           string
	Autor            string
	Idioma           string
	NumPages         int
	Editora          string
	Genero           string
	AnoPublicacao    string   // data de publicação do livro
	LinkCompra       []string // pode ser um ou mais
	Sinopse          string
	ValorMultaDiaria float64
}

func (l *Livro) CalcularMulta(valorOriginal float64) {
	l.ValorMultaDiaria = valorOriginal * 30 / 100
}

type User struct {
	Nome         string
	Email        string
	Senha        string
}

func SetUser(nome, email, senha string) *User {
	user :=  &User{
		Nome: nome,
		Email: email,
		Senha: senha,
	}
	return user
}

// escrever o getter e setter para o atributo senha
