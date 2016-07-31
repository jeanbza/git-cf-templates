package main

import (
    "net/http"
    "fmt"
    "os"

    "github.com/Pallinder/go-randomdata"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        sillyName := randomdata.SillyName()
        w.Write([]byte(fmt.Sprintf("Hello %s\n", sillyName)))
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
