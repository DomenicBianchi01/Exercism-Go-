package luhn

import "strings"
import "strconv"

func Valid(number string) bool {
	number = strings.Replace(number, " ", "", -1)
	if len(number) <= 1 {
		return false
	}

	total := 0

	for i := len(number) - 1; i >= 0; i-- {
		if convertedNumber, err := strconv.Atoi(string(number[i])); err == nil {
			if (len(number)%2 == 0 && i%2 != 0) || len(number)%2 != 0 && i%2 == 0 {
				total += convertedNumber
			} else {
				if convertedNumber*2 > 9 {
					total += (convertedNumber * 2) - 9
				} else {
					total += convertedNumber * 2
				}
			}
		} else {
			return false
		}
	}

	return total%10 == 0
}
