// main.go
package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

// Shape interface for different shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle struct
type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// CreateShapeFactory creates a shape based on type
func CreateShapeFactory(shapeType string, dimensions ...float64) (Shape, error) {
	switch shapeType {
	case "rectangle":
		if len(dimensions) != 2 {
			return nil, errors.New("rectangle requires two dimensions: length and width")
		}
		return Rectangle{Length: dimensions[0], Width: dimensions[1]}, nil
	case "circle":
		if len(dimensions) != 1 {
			return nil, errors.New("circle requires one dimension: radius")
		}
		return Circle{Radius: dimensions[0]}, nil
	case "triangle":
		if len(dimensions) != 3 {
			return nil, errors.New("triangle requires three dimensions: sides A, B, and C")
		}
		if !isValidTriangle(dimensions[0], dimensions[1], dimensions[2]) {
			return nil, errors.New("invalid triangle dimensions")
		}
		return Triangle{A: dimensions[0], B: dimensions[1], C: dimensions[2]}, nil
	default:
		return nil, errors.New("unknown shape type")
	}
}

func isValidTriangle(a, b, c float64) bool {
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	return sides[0]+sides[1] > sides[2]
}

func main() {
	shapes := []struct {
		Type       string
		Dimensions []float64
	}{
		{"rectangle", []float64{10, 5}},
		{"circle", []float64{7}},
		{"triangle", []float64{3, 4, 5}},
	}

	for _, shapeData := range shapes {
		shape, err := CreateShapeFactory(shapeData.Type, shapeData.Dimensions...)
		if err != nil {
			fmt.Printf("Error creating shape: %v\n", err)
			continue
		}
		fmt.Printf("Shape: %s, Area: %.2f, Perimeter: %.2f\n", shapeData.Type, shape.Area(), shape.Perimeter())
	}
}
