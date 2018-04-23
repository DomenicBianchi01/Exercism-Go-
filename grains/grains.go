package grains

import "errors"

func Square(position int) (uint64, error) {
	if position < 1 || position > 64 {
		return 0, errors.New("Position must be greater than 0 and less than 65")
	}

	var total uint64 = 1

	for i := 1; i < position; i++ {
		total *= 2
	}
	return total, nil
}

func Total() uint64 {
	total, _ := Square(64)
	return total*2 - 1
}
