package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (c Triangle) CalcPerimeter() float64 {
	return 3 * c.Side
}

func (c Triangle) CalcArea() float64 {
	return math.Sqrt(3) / 4 * c.Side
}
