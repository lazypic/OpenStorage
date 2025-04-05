package cmd

import (
    "net/http"
    "strings"

    "github.com/gorilla/mux"
    "github.com/lazypic/OpenStorage/cmd/internal"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, []string)

var handlers = map[string]map[string]HandlerFunc{
    "zpool": {
        "list": ZpoolList,
    },
    "zfs": {
        "list": ZfsList,
    },
}

func Dispatch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cmd := vars["cmd"]
    subcmd := vars["subcmd"]
    args := strings.Fields(r.URL.Query().Get("args"))
    allArgs := append([]string{subcmd}, args...)

    if hmap, ok := handlers[cmd]; ok {
        if handler, ok := hmap[subcmd]; ok {
            handler(w, r, allArgs)
            return
        }
    }

    internal.FallbackRaw(w, r, cmd, allArgs)
}
