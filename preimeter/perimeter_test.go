package perimeter

import (
	"testing"
)

func TestPerimeteer(t *testing.T) {

	t.Run("get perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}
func TestArea(t *testing.T) {
	// t.Helper()
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{width: 12, height: 6}, 72.0},
		{"Circle", Circle{Radius: 10}, 314.1592653589793},
		{"Triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.want)
			}
		})

	}
}
