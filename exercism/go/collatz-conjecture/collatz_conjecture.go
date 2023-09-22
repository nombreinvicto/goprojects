package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, errors.New("number has to be greater than 0")
	}

	// initialise the no of steps
	steps := 0

	// start looping
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps += 1
	}
	return steps, nil
}
