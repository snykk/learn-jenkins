// main_test.go
package main

import (
	"math"
	"testing"
)

func TestRectangle(t *testing.T) {
	rect := Rectangle{Length: 10, Width: 5}
	if rect.Area() != 50 {
		t.Errorf("Expected area 50, got %.2f", rect.Area())
	}
	if rect.Perimeter() != 30 {
		t.Errorf("Expected perimeter 30, got %.2f", rect.Perimeter())
	}
}

func TestCircle(t *testing.T) {
	circ := Circle{Radius: 7}
	expectedArea := math.Pi * 7 * 7
	expectedPerimeter := 2 * math.Pi * 7

	// Define a small tolerance for float comparison
	epsilon := 0.0000001

	if math.Abs(circ.Area()-expectedArea) > epsilon {
		t.Errorf("Expected area %.2f, got %.2f", expectedArea, circ.Area())
	}
	if math.Abs(circ.Perimeter()-expectedPerimeter) > epsilon {
		t.Errorf("Expected perimeter %.2f, got %.2f", expectedPerimeter, circ.Perimeter())
	}
}

func TestTriangle(t *testing.T) {
	tri := Triangle{A: 3, B: 4, C: 5}
	expectedArea := 6.0
	expectedPerimeter := 12.0

	if tri.Area() != expectedArea {
		t.Errorf("Expected area %.2f, got %.2f", expectedArea, tri.Area())
	}
	if tri.Perimeter() != expectedPerimeter {
		t.Errorf("Expected perimeter %.2f, got %.2f", expectedPerimeter, tri.Perimeter())
	}
}

func TestCreateShapeFactory(t *testing.T) {
	_, err := CreateShapeFactory("rectangle", 10, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = CreateShapeFactory("circle", 7)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = CreateShapeFactory("triangle", 3, 4, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = CreateShapeFactory("triangle", 1, 2, 10) // Invalid triangle
	if err == nil {
		t.Errorf("Expected error for invalid triangle, got nil")
	}

	_, err = CreateShapeFactory("unknown", 10)
	if err == nil {
		t.Errorf("Expected error for unknown shape, got nil")
	}
}
