package server

import (
	"fmt"
	"net/http"
)

func init() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr: "0.0.0.0:8001",
		Handler: mux,
	}
	mux.HandleFunc("/", handler)
	_ = server.ListenAndServe()
}

func handler(repWriter http.ResponseWriter, req *http.Request) {

}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		_ = fmt.Errorf("error occor in form parse %s", err)
	}

}
