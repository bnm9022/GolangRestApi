package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

var ApplicationMDs []ApplicationMD
var currentId = 0


func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/applicationMDs", returnAllApplicationMDs).Methods("GET")
    router.HandleFunc("/applicationMD/{id}", returnSingleApplicationMD).Methods("GET")
    router.HandleFunc("/applicationMD", returnApplicationMDLists).Methods("GET")
    router.HandleFunc("/applicationMD", createNewApplicationMD).Methods("POST")
    router.HandleFunc("/applicationMD", putApplicationMD).Methods("PUT")
    router.HandleFunc("/applicationMD", deleteApplicationMD).Methods("DELETE")
    http.ListenAndServe(":8000", router)
}