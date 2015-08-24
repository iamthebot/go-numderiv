package numeriv

import (
	"errors"
)

//Calculate a backwards 3rd order approximate 2nd derivative of an int slice
//d: data
//x: index in data
//s: step size
func BackSecondDerivInt(d []int, x int, s int) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3.0*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 3rd order approximate 2nd derivative of an int32 slice
//d: data
//x: index in data
//s: step size
func BackSecondDerivInt32(d []int32, x int32, s int32) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > int32(len(d)) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3.0*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 3rd order approximate 2nd derivative of an int64 slice
//d: data
//x: index in data
//s: step size
func BackSecondDerivInt64(d []int64, x int64, s int64) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > int64(len(d)) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3.0*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 3rd order approximate 2nd derivative of a float32 slice
//d: data
//x: index in data
//s: step size
func BackSecondDerivFloat32(d []float32, x int, s int) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*float64(d[x]) - 5.0*float64(d[x-s]) + 4.0*float64(d[x-2*s]) - float64(d[x-3.0*s])) / (float64(s) * float64(s))), nil
}

//Calculate a backwards 3rd order approximate 2nd derivative of a float64 slice
//d: data
//x: index in data
//s: step size
func BackSecondDerivFloat64(d []float64, x int, s int) (float64, error) {
	if (x - 3*s) < 0 {
		return 0.0, errors.New("step size puts desired point out of bounds")
	} else if x > len(d) || x < 0 {
		return 0.0, errors.New("x is out of bounds")
	} else if s < 0 {
		return 0.0, errors.New("s must be positive")
	}
	return ((2.0*d[x] - 5.0*d[x-s] + 4.0*d[x-2*s] - d[x-3.0*s]) / (float64(s) * float64(s))), nil
}

//Calculate a backwards looking 3rd order approximate 1st derivative of an int slice
//d: data
//x: index in data
//s: step size
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

//Calculate a backwards looking 3rd order approximate 1st derivative of an int32 slice
//d: data
//x: index in data
//s: step size
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

//Calculate a backwards looking 3rd order approximate 1st derivative of an int64 slice
//d: data
//x: index in data
//s: step size
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

//Calculate a backwards looking 3rd order approximate 1st derivative of a float32 slice
//d: data
//x: index in data
//s: step size
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

//Calculate a backwards looking 3rd order approximate 1st derivative of a float64 slice
//d: data
//x: index in data
//s: step size
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
