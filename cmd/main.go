package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/lucas-rabello-dev/WEB-CRUD.git/server"

	"github.com/joho/godotenv"
)

const (
	HomeWay string = "./templates/home/home.html"
	CrudPageWay string = "./templates/crudPage/crudPage.html"
)

func main() {
	// Load .env
	port, host, err := LoadDotEnv()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	mux = server.ServeStaticFiles(mux)

	mux.HandleFunc("/home/", HandleHome)

	mux.HandleFunc("/home/crudPage/", HandleCrud)


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
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "método inválido | método é diferente de GET", http.StatusMethodNotAllowed)
	}

	t, err := template.ParseFiles(HomeWay)
	if err != nil {
		http.Error(w, "erro ao carregar o html", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "home", nil)
}

func HandleCrud(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "método inválido | método é diferente de GET", http.StatusMethodNotAllowed)
	}

	t, err := template.ParseFiles(CrudPageWay)
	if err != nil {
		http.Error(w, "erro ao carregar o html", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "crudPage", nil)
}