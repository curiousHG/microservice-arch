package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, from service A!")
    log.Println("Hello, from service A!")
    // log.Fatal("Hello, from service A!")
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    message := vars["message"]
    fmt.Fprintf(w, "Message: %s", message)
    log.Println("Message: %s", message)
}

func main() {
    port := "8080"
    r := mux.NewRouter()
    r.HandleFunc("/", helloHandler).Methods("GET")
    r.HandleFunc("/msg/{message}", messageHandler).Methods("GET")

    http.Handle("/", r)
    http.ListenAndServe(":" + port, nil)

}
