package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
type tableTests struct {
	name string
	shape Shape
	want float64
}
func TestArea(t *testing.T) {

	var areaTests []tableTests = []tableTests{
		{name : "test1", shape: Rectangle{2,4}, want : 8},
		{name : "test2", shape: Rectangle{4,5}, want : 20},
		{name : "test3", shape: triangle{4,5}, want : 10},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
