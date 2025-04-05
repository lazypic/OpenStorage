package cmd

import (
    "fmt"
    "net/http"
    "github.com/lazypic/OpenStorage/parser"
    "github.com/lazypic/OpenStorage/util"
)

func DfList(w http.ResponseWriter, r *http.Request, args []string) {
    fmt.Println(args)
    out, err := util.RunCommand("df", args)
    if err != nil {
        util.RespondError(w, err, out)
        return
    }
    data := parser.ParseDfList(out)
    util.RespondJSON(w, data)
}
