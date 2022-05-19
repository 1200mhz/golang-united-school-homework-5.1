package square

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEnd(t *testing.T) {
	square := &Square{
		Point{
			10,
			20,
		},
		30,
	}

	actual := square.End()
	expected := Point{40, -10}

	if !reflect.DeepEqual(actual, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}

func TestArea(t *testing.T) {
	square := &Square{
		Point{
			10,
			20,
		},
		30,
	}

	actual := square.Area()
	expected := uint(900)

	if !reflect.DeepEqual(actual, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}

func TestPerimeter(t *testing.T) {
	square := &Square{
		Point{
			10,
			20,
		},
		30,
	}

	actual := square.Perimeter()
	expected := uint(120)

	if !reflect.DeepEqual(actual, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}
