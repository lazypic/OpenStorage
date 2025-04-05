package main

import (
    "github.com/gorilla/mux"
    "github.com/lazypic/OpenStorage/cmd"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/run/{cmd}/{subcmd}", cmd.Dispatch).Methods("GET")
    return r
}
