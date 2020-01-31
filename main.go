package main

import (
    "fmt"
    "net/http"
    "os"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, "Hello, word!")
}

func main() {
    port := os.Getenv("PORT") // Heroku provides the port to bind to
    if port == "" {
      port = "8080"
    }
    http.HandleFunc("/", helloServer)
    http.ListenAndServe(":" + port, nil)
}
