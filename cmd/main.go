package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/lucas-rabello-dev/WEB-CRUD.git/server"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

const (
	HomeWay string = "./templates/home/home.html"
	CrudPageWay string = "./templates/cadastro-livro/cadastro-livro.html"
	LoginWay string = "./templates/login/login.html"
	MultasWay string = "./templates/multas/multas.html"
	EstoqueWay string = "./templates/estoque/estoque.html"
	EmprestimosWay string = "./templates/emprestimos/emprestimos.html"
	BloqueadosWay string = "./templates/bloqueados/bloqueados.html"
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

	mux.HandleFunc("/home/", HandleHome)
	success.Println("'home' carregada ✅")

	mux.HandleFunc("/home/cadastro-livro/", HandleCadastroLivro)
	success.Println("'cadastro-livro' carregado ✅")

	mux.HandleFunc("/", HandleLogin)
	success.Println("'login' carregado ✅")

	mux.HandleFunc("/home/multas/", HandleMultas)
	success.Println("'multas' carregado ✅")

	mux.HandleFunc("/home/estoque/", HandleEstoque)
	success.Println("'estoque' carregado ✅")

	mux.HandleFunc("/home/emprestimos/", HandleEmprestimos)
	success.Println("'emprestimos' carregado ✅")

	mux.HandleFunc("/home/bloqueados/", HandleBloqueados)
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

func HandleHome(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home", "/home/":
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "método inválido | método é diferente de GET", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles(HomeWay)
		if err != nil {
			http.Error(w, "erro ao carregar o html", http.StatusInternalServerError)
			err_.Println("Erro ao carregar html em:", r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "home", nil)
	default:
		http.NotFound(w, r)
	}
}

func HandleCadastroLivro(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home/cadastro-livro", "/home/cadastro-livro/":
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "método inválido | método é diferente de GET", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles(CrudPageWay)
		if err != nil {
			http.Error(w, "erro ao carregar o html", http.StatusInternalServerError)
			err_.Println("Erro ao carregar html em:", r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "cadastro-livro", nil)
	default:
		http.NotFound(w, r)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "método inválido | método diferente de GET", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles(LoginWay)
		if err != nil {
			http.Error(w, "erro ao carregar html", http.StatusInternalServerError)
			err_.Println("Erro ao carregar html em:", r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "login", nil)
	default:
		http.NotFound(w, r)
	}
}

func HandleMultas(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/home/multas", "/home/multas/":
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "método inválido | método diferente de GET", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles(MultasWay)
		if err != nil {
			http.Error(w, "erro ao carregar html", http.StatusInternalServerError)
			err_.Println("Erro ao carregar html em:", r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "multas", nil)
    default:
        http.NotFound(w, r)
    }
}

func HandleEstoque(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home/estoque", "/home/estoque/":
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "método inválido | método diferente de GET", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles(EstoqueWay)
		if err != nil {
			http.Error(w, "erro ao carregar html", http.StatusInternalServerError)
			err_.Println("Erro ao carregar html em:", r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "estoque", nil)
	default:
		http.NotFound(w, r)
	}
}

func HandleEmprestimos(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home/emprestimos", "/home/emprestimos/":
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "método inválido | método diferente de GET", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles(EmprestimosWay)
		if err != nil {
			http.Error(w, "erro ao carregar html", http.StatusInternalServerError)
			err_.Println("Erro ao carregar html em:", r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "emprestimos", nil)
	default:
		http.NotFound(w, r)
	}
}

func HandleBloqueados(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home/bloqueados", "/home/bloqueados/":
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "método inválido | método diferente de GET", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles(BloqueadosWay)
		if err != nil {
			http.Error(w, "erro ao carregar html", http.StatusInternalServerError)
			err_.Println("Erro ao carregar html em:", r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "bloqueados", nil)
	default:
		http.NotFound(w, r)
	}
}