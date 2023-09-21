package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC]|\[DBG]|\[INF]|\[WRN]|\[ERR]|\[FTL])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`(<\**>|<~*>|<-*>|<=>|<-\*~\*->)`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	sliceOfLogs := []string{}
	for _, line := range lines {
		re := regexp.MustCompile(`User\s+`)

		if re.MatchString(line) {
			secondPart := re.Split(line, -1)[1]
			username := strings.Split(secondPart, " ")[0]
			log := fmt.Sprintf("[USR] %s %s", username, line)
			sliceOfLogs = append(sliceOfLogs, log)
		} else {
			sliceOfLogs = append(sliceOfLogs, line)
		}
	}
	return sliceOfLogs
}
