package main

import (
        "fmt"
        "net/http"
        "runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Go version: %s", runtime.Version())
}

func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8080", nil)
}