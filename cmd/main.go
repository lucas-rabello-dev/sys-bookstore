package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lucas-rabello-dev/WEB-CRUD.git/server"
	"github.com/lucas-rabello-dev/WEB-CRUD.git/handlers"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)


var (
	success *color.Color = color.New(color.FgGreen).Add(color.Bold, color.Italic)
	err_ *color.Color = color.New(color.FgRed).Add(color.Bold, color.Italic)
)

func main() {
	// Load .env
	port, host, err := LoadDotEnv()
	if err != nil {
		// log.Fatal(err)
		err_.Printf("Error: %v ❌\n", err)
		return
	}
	success.Println(".env carregado ✅")

	mux := http.NewServeMux()

	// servindo a pasta static
	mux = server.ServeStaticFiles(mux)

	mux.HandleFunc("/home/", handlers.HandleHome)
	success.Println("'home' carregada ✅")

	mux.HandleFunc("/home/cadastro-livro/", handlers.HandleCadastroLivro)
	success.Println("'cadastro-livro' carregado ✅")

	mux.HandleFunc("/", handlers.HandleLogin)
	success.Println("'login' carregado ✅")

	mux.HandleFunc("/home/multas/", handlers.HandleMultas)
	success.Println("'multas' carregado ✅")

	mux.HandleFunc("/home/estoque/", handlers.HandleEstoque)
	success.Println("'estoque' carregado ✅")

	mux.HandleFunc("/home/emprestimos/", handlers.HandleEmprestimos)
	success.Println("'emprestimos' carregado ✅")

	mux.HandleFunc("/home/bloqueados/", handlers.HandleBloqueados)
	success.Println("'bloqueados' carregado ✅")

	fmt.Printf("Rodando em: http://%s:%s\n", host, port)
	http.ListenAndServe(":" + port, mux)
}

func LoadDotEnv() (string, string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", "", fmt.Errorf("houve um erro ao carregar o .env | tente rodar com 'go run cmd/main.go'")
	}

	port_env, ok:= os.LookupEnv("PORTA")
	if !ok {
		return "", "", fmt.Errorf("impossível carregar a variavel 'PORTA' do .env")
	}

	host_env, ok := os.LookupEnv("HOST")
	if !ok {
		return "", "", fmt.Errorf("impossível carregar a variável 'HOST' do .env")
	}

	return port_env, host_env, nil
}
