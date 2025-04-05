package cmd

import (
    "net/http"
    "github.com/lazypic/OpenStorage/parser"
    "github.com/lazypic/OpenStorage/util"
)

func ZfsList(w http.ResponseWriter, r *http.Request, args []string) {
    out, err := util.RunCommand("zfs", args)
    if err != nil {
        util.RespondError(w, err, out)
        return
    }
    data := parser.ParseZfsList(out)
    util.RespondJSON(w, data)
}
