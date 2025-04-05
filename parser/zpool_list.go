package parser

import "strings"

func ParseZpoolList(output string) []map[string]string {
    lines := strings.Split(output, "\n")
    var result []map[string]string

    for _, line := range lines {
        if line == "" {
            continue
        }
        cols := strings.Fields(line)
        if len(cols) >= 10 {
            result = append(result, map[string]string{
                "name":     cols[0],
                "size":     cols[1],
                "alloc":    cols[2],
                "free":     cols[3],
                "expandsz": cols[4],
                "frag":     cols[5],
                "cap":      cols[6],
                "dedup":    cols[7],
                "health":   cols[8],
                "altroot":  cols[9],
            })
        }
    }
    return result
}