package core

import (
    "fmt"
    "regexp"
    "time"
)

var lineStartRegex = regexp.MustCompile(`(?m)^`)

func formatDuration(d time.Duration) string {
    str := ""
    sec := uint64(d / time.Second)
    if (sec / 31536000) != 0 {
        str += fmt.Sprintf("%dy", (sec / 31536000))
    }
    if (sec / 86400) % 365 != 0 {
        str += fmt.Sprintf("%dd", (sec / 86400) % 365)
    }
    if (sec / 3600) % 24 != 0 {
        str += fmt.Sprintf("%dh", (sec / 3600) % 24)
    }
    if (sec / 60) % 60 != 0 {
        str += fmt.Sprintf("%dm", (sec / 60) % 60)
    }
    if sec % 60 != 0 {
        str += fmt.Sprintf("%ds", sec % 60)
    }
    return str
}
