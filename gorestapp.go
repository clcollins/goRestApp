package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"html/template"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"github.com/ajays20078/go-http-logger"

	. "github.com/clcollins/goRestApp/models"
	. "github.com/clcollins/goRestApp/dao"
)

var dao = GatosDAO{}

func CreateCat(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gato Gato
	if err := json.NewDecoder(r.Body).Decode(&gato); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	gato.ID = bson.NewObjectId()
	if err := dao.Insert(gato); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, gato)
}

func ReadCat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gato, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Gato ID")
		return
	}
	respondWithJson(w, http.StatusOK, gato)
}

func UpdateCat(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gato Gato
	if err := json.NewDecoder(r.Body).Decode(&gato); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := dao.Update(gato); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteCat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "We don't do that, here.  We're a no-kill API.")
}

func CatParade(w http.ResponseWriter, r *http.Request) {
	gatos, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, gatos)
}

func CatParadeTmpl(w http.ResponseWriter, r *http.Request) {
	gatos, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	t := template.New("Cat Parade")
	t, _  = t.Parse(`Cat Parade!
		  {{ with .Gatos }}
			  {{ range .}}
				    {{ .Image }}
				{{ end }}
			{{ end }}
  `)
		t.Execute(w, gatos)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init(){
	dao.Server = "db"
	dao.Database = "gatos"
	dao.Connect()
}

func main() {
  listenPort := "3000"
	r          := mux.NewRouter()
	apiV1      := "/api/v1"

	r.HandleFunc("/cats", CatParadeTmpl).Methods("GET")

	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), CatParade).Methods("GET")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), CreateCat).Methods("POST")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), UpdateCat).Methods("PUT")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), UpdateCat).Methods("PATCH")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), DeleteCat).Methods("DELETE")
	r. HandleFunc(fmt.Sprintf("%s/gatos/{id}", apiV1), ReadCat).Methods("GET")

  fmt.Printf("Listening on :" + listenPort + "\n")
	if err := http.ListenAndServe(":" + listenPort , httpLogger.WriteLog(r, os.Stdout))
	err != nil { log.Fatal(err)	}
}
