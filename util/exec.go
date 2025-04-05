package util

import (
    "encoding/json"
    "net/http"
    "os/exec"
)

func RunCommand(name string, args []string) (string, error) {
    out, err := exec.Command(name, args...).CombinedOutput()
    return string(out), err
}

func RespondJSON(w http.ResponseWriter, v any) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(v)
}

func RespondError(w http.ResponseWriter, err error, raw string) {
    w.WriteHeader(http.StatusInternalServerError)
    RespondJSON(w, map[string]string{
        "error":  err.Error(),
        "output": raw,
    })
}