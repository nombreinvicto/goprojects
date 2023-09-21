package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := []string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long."}

	for _, line := range lines {
		re := regexp.MustCompile(`User\s+`)

		if re.MatchString(line) {
			secondPart := re.Split(line, -1)[1]
			username := strings.Split(secondPart, " ")[0]
			fmt.Println(username)

		}

	}
}
