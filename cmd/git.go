package cmd

import (
    "fmt"
    "net/http"
    "github.com/lazypic/OpenStorage/parser"
    "github.com/lazypic/OpenStorage/util"
)

func GitVersion(w http.ResponseWriter, r *http.Request, args []string) {
    fmt.Println(args)
    out, err := util.RunCommand("git", args)
    if err != nil {
        util.RespondError(w, err, out)
        return
    }
    data := parser.ParseGitVersion(out)
    util.RespondJSON(w, data)
}
