package vec3

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := Vec3[float32]{X: 1, Y: 2, Z: 3}
	b := Vec3[float32]{X: 10, Y: 11, Z: 12}

	res := Add(&a, &b)
	if !Equals(&res, &Vec3[float32]{11, 13, 15}) {
		t.Errorf("Expected {11, 13, 15}, got %v", res)
	}
}
