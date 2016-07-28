package main

import (
    "net/http"
    "fmt"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello world"))
    })

    var port string
    if os.Getenv("PORT") == "" {
        port = "8080"
    } else {
        port = os.Getenv("PORT")
    }

    fmt.Printf("Listening on %s\n", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
        panic(err)
    }
}
