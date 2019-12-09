package server

import (
	"../model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var appAddress string = "0.0.0.0:8001"

func init() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    appAddress,
		Handler: mux,
	}
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/sickness", sicknessHandler)
	_ = server.ListenAndServe()
	log.Printf("server is running at: %s", appAddress)
}

func handler(repWriter http.ResponseWriter, req *http.Request) {
	log.Printf("resolve request from %s", req.URL)
}

func sicknessHandler(repWriter http.ResponseWriter, req *http.Request){
	switch req.Method {
	case "GET":
		sicknessFetchHandler(repWriter, req)
		return
	case "POST":
		sicknessSaveHandler(repWriter, req)
		return
	case "DELETE":

	case "PUT":

	}
}

func sicknessSaveHandler(repWriter http.ResponseWriter, req *http.Request) {
	l := req.ContentLength
	body := make([]byte, l)
	_, err := req.Body.Read(body)
	if err != nil {
		log.Println("reading request body error .", err)
	}
	var sick model.Sickness
	err = json.Unmarshal(body, &sick)
	if err != nil {
		log.Println("parsing request json body error .", err)
	}
	err = sick.Save()
	if err != nil {
		log.Println("saving sick info occur error. ", err)
	}else {
		log.Println("saved sickness info,", sick)
	}
	_ = writeSicknessToResponse(repWriter, &sick)
}

func sicknessFetchHandler(repWriter http.ResponseWriter, req *http.Request) {

	sick, err := parseAndFetchSickness(req)
	if err != nil {
		log.Println("find sick info error. ", err)
		return
	}
	err = writeSicknessToResponse(repWriter, sick)
	if err != nil {
		log.Println("error occurred on write response")
	}
}

func parseAndFetchSickness(req *http.Request) (sick *model.Sickness, err error) {
	err = req.ParseForm()
	if err != nil {
		return
	}
	form := req.Form
	id := form.Get("id")
	var sickId int64
	sickId, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	sick, err = model.FetchSickById(sickId)
	return
}

func writeSicknessToResponse(repWriter http.ResponseWriter, sick *model.Sickness) (err error) {
	var btys []byte
	btys, err = writeSicknessToJson(sick)
	if err != nil {
		log.Println("write sickness to json error.")
		return err
	}
	_, err = repWriter.Write(btys)
	if err != nil {
		log.Println("write sickness json to http response error.")
		return err
	}
	return nil
}

func writeSicknessToJson(sick *model.Sickness) (res []byte, err error) {
	res, err = json.Marshal(sick)
	return
}

func writeSicknessListToJson(sicks *[]model.Sickness) (res []byte, err error) {
	res, err = json.Marshal(sicks)
	return
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		_ = fmt.Errorf("error occor in form parse %s", err)
	}

}
