package bob

import "strings"
import "regexp"

func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	alphaRegex := regexp.MustCompile("[a-zA-Z]")

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	uppercaseRemark := strings.ToUpper(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	containsLetters := alphaRegex.MatchString(remark)

	if remark == uppercaseRemark && containsLetters && !isQuestion {
		return "Whoa, chill out!"
	} else if remark == uppercaseRemark && containsLetters {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
