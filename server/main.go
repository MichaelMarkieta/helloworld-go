package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func handlerfoo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "foo")
}

func handlerbar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "bar")
}

func handlerfoobar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "foobar")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/foo", handlerfoo)
    http.HandleFunc("/bar", handlerbar)
    http.HandleFunc("/foobar", handlerfoobar)
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
