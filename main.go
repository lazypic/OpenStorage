package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

type ZpoolDevice struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Read  string `json:"read"`
	Write string `json:"write"`
	Cksum string `json:"cksum"`
}

type ZpoolStatus struct {
	PoolName  string        `json:"pool_name"`
	Status    string        `json:"status"`
	Scan      string        `json:"scan,omitempty"`
	Errors    string        `json:"errors"`
	Devices   []ZpoolDevice `json:"devices"`
	RawOutput string        `json:"raw_output"`
}

func parseZpoolStatus(output string) ZpoolStatus {
	lines := strings.Split(output, "\n")
	var pool ZpoolStatus
	pool.RawOutput = output
	var inDeviceSection bool

	// 정규식: 디바이스 라인 패턴
	deviceLineRegex := regexp.MustCompile(`^\s*(\S+)\s+(\S+)\s+(\d+)\s+(\d+)\s+(\d+)`)

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
			inDeviceSection = false
		} else if strings.HasPrefix(line, "config:") {
			inDeviceSection = true
		} else if inDeviceSection {
			if matches := deviceLineRegex.FindStringSubmatch(line); matches != nil {
				device := ZpoolDevice{
					Name:  matches[1],
					State: matches[2],
					Read:  matches[3],
					Write: matches[4],
					Cksum: matches[5],
				}
				pool.Devices = append(pool.Devices, device)
			}
		}
	}

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
	fmt.Println("Listening on :9090")
	http.ListenAndServe(":9090", nil)
}
