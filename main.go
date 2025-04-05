package main

import (
    "log"
    "net/http"
)

func main() {
    router := SetupRouter()
    log.Println("OpenStorag API listening on :9090")
    log.Fatal(http.ListenAndServe(":9090", router))
}
