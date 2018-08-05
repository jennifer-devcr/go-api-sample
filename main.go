package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Person struct {
  ID string `json:"id,omitempty"`
  FirstName string `json:"firstname,omitempty"`
  LastName string `json:"lastname,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func main() {
  router := mux.NewRouter()

  people = append(people, Person{ID: "1", FirstName: "Raul", LastName: "Lopez"})

  router.HandleFunc("/people", GetPeople).Methods("GET")

  log.Fatal(http.ListenAndServe(":3000", router))
}
