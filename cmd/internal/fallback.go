package internal

import (
    "net/http"
    "github.com/lazypic/OpenStorage/util"
)

func FallbackRaw(w http.ResponseWriter, r *http.Request, cmd string, args []string) {
    out, err := util.RunCommand(cmd, args)
    if err != nil {
        util.RespondError(w, err, out)
        return
    }
    util.RespondJSON(w, map[string]string{
        "output": out,
    })
}
