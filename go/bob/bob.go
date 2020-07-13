// Package bob provides responses to conversational input
package bob

import (
	"regexp"
	"strings"
)

// Hey accepts a remark and prints a lackadaisical response
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	hasAlpha, _ := regexp.MatchString(`[A-Za-z]`, trimmedRemark)
	isAllUpperCase := hasAlpha && strings.ToUpper(trimmedRemark) == trimmedRemark
	isQuestion := strings.HasSuffix(trimmedRemark, "?")
	isExclamation := strings.HasSuffix(trimmedRemark, "!")

	if !isAllUpperCase && isQuestion {
		return "Sure."
	} else if (isAllUpperCase && isExclamation) || (isAllUpperCase && !isExclamation && !isQuestion) {
		return "Whoa, chill out!"
	} else if isAllUpperCase && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else {
		return "Whatever."
	}
}
