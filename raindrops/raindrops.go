package raindrops

import "strconv"

func Convert(drops int) string {
	sounds := ""

	if drops%3 == 0 {
		sounds += "Pling"
	}
	if drops%5 == 0 {
		sounds += "Plang"
	}
	if drops%7 == 0 {
		sounds += "Plong"
	}

	if len(sounds) == 0 {
		return strconv.Itoa(drops)
	}

	return sounds
}
