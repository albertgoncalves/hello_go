package main

// via https://github.com/gorilla/mux

import (
    "log"
    "mux-1.6.2"
    "net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", YourHandler)
    log.Fatal(http.ListenAndServe(":8000", r))
}
