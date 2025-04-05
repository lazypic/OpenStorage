package parser

import "strings"
import "fmt"

func ParseGitVersion(output string) []map[string]string {
    lines := strings.Split(output, "\n")
    var result []map[string]string

    for _, line := range lines {
        if line == "" {
            continue
        }
        cols := strings.Fields(line)
        if len(cols) >= 10 {
            result = append(result, map[string]string{
                "app":     cols[0],
                "string":     cols[1],
                "version":    cols[2],
            })
        }
    }
    return result
}