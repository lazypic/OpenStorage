package parser

import "strings"
import "fmt"

func ParseDfList(output string) []map[string]string {
    lines := strings.Split(output, "\n")
    var result []map[string]string

    for _, line := range lines {
        if line == "" {
            continue
        }
        cols := strings.Fields(line)
        if len(cols) >= 10 {
            result = append(result, map[string]string{
                "filesystem":     cols[0],
                "size":     cols[1],
                "use":    cols[2],
                "free":     cols[3],
                "percent": cols[4],
                "mount":     cols[5],
            })
        }
    }
    fmt.Println("test")
    return result
}