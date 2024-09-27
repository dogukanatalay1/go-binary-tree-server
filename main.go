package main

import (
    "log"
    "net/http"

	handlers "golang-binary-tree/handlers"
)

func main() {
    http.HandleFunc("/maxPathSum", handlers.MaxPathSumHandler)
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
