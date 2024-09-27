package handlers

import (
    "net/http"
    "io"
    "log"
)

func MaxPathSumHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("Error reading body: %v", err)
        return
    }
    log.Printf("Received body: %s", body)
}


