package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of boxw
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	//panic("implement me")
	if cap(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return errors.New("Ошибка")
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	//	panic("implement me")

	if i < len(b.shapes) {
		return b.shapes[i], nil
	} else {
		return nil, errors.New("Error")
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	//	panic("implement me")

	if i < len(b.shapes) {
		mShape := b.shapes[i]
		b.shapes[i] = b.shapes[len(b.shapes)-1]
		b.shapes[len(b.shapes)-1] = nil
		b.shapes = b.shapes[:len(b.shapes)-1]
		return mShape, nil
	} else {
		return nil, errors.New("ошибка")

	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	//panic("implement me")
	if i < len(b.shapes) {
		mShape := b.shapes[i]
		b.shapes[i] = shape
		return mShape, nil
	} else {
		return nil, errors.New("ошибка")

	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	//	panic("implement me")
	sum := 0.00

	for _, elem := range b.shapes {
		//		switch elem.(type) {
		//		case Circle:
		//			sum = sum + elem.CalcPerimeter()
		//		case Triangle:
		//			sum = sum + elem.CalcPerimeter()
		//		case Rectangle:
		//			sum = sum + elem.CalcPerimeter()
		//		default:
		//			sum = sum + 0.00
		//		}
		sum = sum + elem.CalcPerimeter()

	}

	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	//panic("implement me")
	sum := 0.00
	for _, elem := range b.shapes {
		//		switch elem.(type) {
		//		case Circle:
		//			sum = sum + elem.CalcArea()
		//		case Triangle:
		//			sum = sum + elem.CalcArea()
		//		case Rectangle:
		//			sum = sum + elem.CalcArea()
		//		default:
		//			sum = sum + 0.00
		sum = sum + elem.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	//	panic("implement me")
	//for i, elem := range b.shapes {
	haveCircle := false
	//j := 0
	newSlice := []Shape{}
	for _, elem := range b.shapes {
		switch elem.(type) {
		case Circle:
			haveCircle = true
		//	b.shapes[i] = b.shapes[len(b.shapes)-1]
		//	j = j + 1
		//	b.shapes[len(b.shapes)-j] = nil

		default:
			newSlice = append(newSlice, elem)
		}
	}
	if haveCircle {
		b.shapes = newSlice
		return nil
	} else {
		return errors.New("jib,rf")
	}

	//}

	return nil
	//if !haveCircle {
	//		return errors.New("ошибка")
	//} else {
	//	b.shapes = b.shapes[:len(b.shapes)-j]
	//	return nil
	//}

}
