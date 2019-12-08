package server

import (
	"../model"
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

var app_address string = "0.0.0.0:8001"

func init() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr: app_address,
		Handler: mux,
	}
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/sickness", sicknessHandler)
	_ = server.ListenAndServe()
	log.Printf("server is running at: %s", app_address)
}

func handler(repWriter http.ResponseWriter, req *http.Request) {
	log.Printf("resolve request from %s", req.URL)
}

func sicknessHandler(repWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		sick := model.Sickness{}
		byt,_ := json.Marshal(&sick)
		repWriter.Write(byt)
		return;
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		_ = fmt.Errorf("error occor in form parse %s", err)
	}

}
