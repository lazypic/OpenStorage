package cmd

import (
    "net/http"
    "github.com/lazypic/OpenStorage/parser"
    "github.com/lazypic/OpenStorage/util"
)

func ZpoolList(w http.ResponseWriter, r *http.Request, args []string) {
    out, err := util.RunCommand("zpool", args)
    if err != nil {
        util.RespondError(w, err, out)
        return
    }
    data := parser.ParseZpoolList(out)
    util.RespondJSON(w, data)
}
