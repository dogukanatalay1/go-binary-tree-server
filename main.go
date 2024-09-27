package main

import (
    "log"
    "net/http"

    handlers "golang-binary-tree/handlers"
)

func main() {
    http.HandleFunc("/max-path-sum", handlers.HandleMaxPathSum)
    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
