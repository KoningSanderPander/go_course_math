package geometry

import "errors"

func Cube(x int) (int, error) {
	if x == 0 {
		return 0, errors.New("cubes with edge length 0 are not allowed")
	}

	return x * x * x, nil
}
