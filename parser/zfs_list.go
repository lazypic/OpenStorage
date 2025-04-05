package parser

import "strings"

func ParseZfsList(output string) []map[string]string {
    lines := strings.Split(output, "\n")
    var result []map[string]string

    for _, line := range lines {
        if line == "" {
            continue
        }
        cols := strings.Fields(line)
        if len(cols) >= 5 {
            result = append(result, map[string]string{
                "name":        cols[0],
                "used":        cols[1],
                "avail":       cols[2],
                "referenced":  cols[3],
                "mountpoint":  cols[4],
            })
        }
    }
    return result
}
