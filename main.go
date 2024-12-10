package main

import (
	"fmt"
	"net/http"

	"github.com/najq/go-server-guess-game/handlers"
)



func main() {
    http.HandleFunc("/hello", handlers.HelloHandler)

    
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}