package utils

import (
	"log"
	"regexp"
	"rss-reader/globals"
	"strings"
)

func MatchStr(str string, callback func(string)) {
	strFinal := strings.TrimSpace(str)
	for _, pattern := range globals.MatchList {
		pattern = "(?i)" + pattern
		re, err := regexp.Compile(pattern)
		if err != nil {
			log.Printf("⚠️ Invalid regular expression: %s, error: %v", pattern, err)
			continue
		}

		if re.MatchString(strFinal) {
			callback(str)
			return
		}
	}
}
