package acronym

import "strings"

func Abbreviate(s string) string {
	dashedString := strings.Replace(s, " ", "-", -1)
	splitString := strings.Split(dashedString, "-")

	abbreviation := ""

	for _, element := range splitString {
		abbreviation += element[:1]
	}

	return strings.ToUpper(abbreviation)
}
