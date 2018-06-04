package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func ReadCat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func UpdateCat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func DeleteCat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "We don't do that, here.  We're a no-kill API.")
}

func CatParade(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func main() {
	r     := mux.NewRouter()
	apiV1 := "/api/v1"

	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), CatParade).Methods("GET")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), CreateCat).Methods("POST")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), UpdateCat).Methods("PUT")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), UpdateCat).Methods("PATCH")
	r. HandleFunc(fmt.Sprintf("%s/gatos", apiV1), DeleteCat).Methods("DELETE")
	r. HandleFunc(fmt.Sprintf("%s/gatos/{id}", apiV1), ReadCat).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
