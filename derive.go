package numeriv

import (
	"errors"
)

/*
* NumDeriv implements functions for simple differentiation of uniformly sampled time-series
* Provided are first and second order approximate derivative calculations using the finite-difference method
 */

//Calculate a backwards 2nd order approximate 2nd derivative of an int slice
func BackSecondDerivInt(d []int, x int, s int) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 2nd order approximate 2nd derivative of an int32 slice
func BackSecondDerivInt32(d []int32, x int32, s int32) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > int32(len(d)) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 2nd order approximate 2nd derivative of an int64 slice
func BackSecondDerivInt64(d []int64, x int64, s int64) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > int64(len(d)) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 2nd order approximate 2nd derivative of a float32 slice
func BackSecondDerivFloat32(d []float32, x int, s int) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 2nd order approximate 2nd derivative of a float64 slice
func BackSecondDerivFloat64(d []float64, x int, s int) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*d[x] - 5.0*d[x-s] + 4.0*d[x-2*s] - d[x-3*s]) / (float64(s) * float64(s))), nil
}

//Calculate a backwards looking 2nd order approximate 1st derivative of an int slice
func BackFirstDerivInt(d []int, x int, s int) (float64, error) {
	if (x - 2*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((3.0*float64(d[x]) - 4.0*float64(d[x-s]) - float64(d[x-2*s])) / (2.0 * float64(s))), nil
}

//Calculate a backwards looking 2nd order approximate 1st derivative of an int32 slice
func BackFirstDerivInt32(d []int32, x int32, s int32) (float64, error) {
	if (x - 2*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > int32(len(d)) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((3.0*float64(d[x]) - 4.0*float64(d[x-s]) - float64(d[x-2*s])) / (2.0 * float64(s))), nil
}

//Calculate a backwards looking 2nd order approximate 1st derivative of an int64 slice
func BackFirstDerivInt64(d []int64, x int64, s int64) (float64, error) {
	if (x - 2*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > int64(len(d)) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((3.0*float64(d[x]) - 4.0*float64(d[x-s]) - float64(d[x-2*s])) / (2.0 * float64(s))), nil
}

//Calculate a backwards looking 2nd order approximate 1st derivative of a float32 slice
func BackFirstDerivFloat32(d []float32, x int, s int) (float64, error) {
	if (x - 2*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((3.0*float64(d[x]) - 4.0*float64(d[x-s]) - float64(d[x-2*s])) / (2.0 * float64(s))), nil
}

//Calculate a backwards looking 2nd order approximate 1st derivative of a float64 slice
func BackFirstDerivFloat64(d []float64, x int, s int) (float64, error) {
	if (x - 2*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return (3.0*d[x] - 4.0*d[x-s] - d[x-2*s]) / (2.0 * float64(s)), nil
}
