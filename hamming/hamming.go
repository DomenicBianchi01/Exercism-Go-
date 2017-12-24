package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Invalid length")
	}

	count := 0

	for index, letter := range a {
		if string(letter) != string(b[index]) {
			count += 1
		}
	}

	return count, nil
}
