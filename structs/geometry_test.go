package structs

import "testing"

func assertEquals(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestParameter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assertEquals(t, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 20.0}, want: 200},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, row := range areaTests {
		t.Run(row.name, func(t *testing.T) {
			got := row.shape.Area()
			if got != row.want {
				t.Errorf("#%v got %g want %g", row.name, got, row.want)
			}
		})
	}
}
