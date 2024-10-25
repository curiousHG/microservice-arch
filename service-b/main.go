package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, from service B!")
    log.Println("Hello, from service B!")
    // log.Fatal("Hello, from service A!")
}

func main() {
    port := "8080"
    r := mux.NewRouter()
    r.HandleFunc("/", helloHandler).Methods("GET")
    http.Handle("/", r)
    http.ListenAndServe(":" + port, nil)



}
