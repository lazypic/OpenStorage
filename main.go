package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type ZpoolStatus struct {
	PoolName  string   `json:"pool_name"`
	Status    string   `json:"status"`
	Scan      string   `json:"scan,omitempty"`
	Errors    string   `json:"errors"`
	Devices   []string `json:"devices"`
	RawOutput string   `json:"raw_output"`
}

func parseZpoolStatus(output string) ZpoolStatus {
	lines := strings.Split(output, "\n")
	var pool ZpoolStatus
	pool.RawOutput = output
	var devices []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "pool:") {
			pool.PoolName = strings.TrimSpace(strings.TrimPrefix(line, "pool:"))
		} else if strings.HasPrefix(line, "state:") {
			pool.Status = strings.TrimSpace(strings.TrimPrefix(line, "state:"))
		} else if strings.HasPrefix(line, "scan:") {
			pool.Scan = strings.TrimSpace(strings.TrimPrefix(line, "scan:"))
		} else if strings.HasPrefix(line, "errors:") {
			pool.Errors = strings.TrimSpace(strings.TrimPrefix(line, "errors:"))
		} else if strings.HasPrefix(line, "mirror") || strings.HasPrefix(line, "raidz") || strings.HasPrefix(line, "NAME") {
			// skip headers or raidz/mirror labels
			continue
		} else if strings.Contains(line, "ONLINE") || strings.Contains(line, "DEGRADED") || strings.Contains(line, "FAULTED") {
			devices = append(devices, line)
		}
	}
	pool.Devices = devices
	return pool
}

func zpoolStatusHandler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("zpool", "status").Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to run zpool status: %v", err), http.StatusInternalServerError)
		return
	}
	status := parseZpoolStatus(string(out))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func main() {
	http.HandleFunc("/", zpoolStatusHandler)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
