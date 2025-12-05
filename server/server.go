package server

import "net/http"

func ServeStaticFiles(mux *http.ServeMux) *http.ServeMux {
	staticHandler := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	return mux
}
