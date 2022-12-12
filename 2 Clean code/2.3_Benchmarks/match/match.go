package match

import (
	"regexp"
	"strings"
)

func MatchContains(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

func MatchRegexp(pattern string, src string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(src)
}
