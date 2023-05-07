package vec3

import (
	"math"
	"testing"
)

type TestVec = Vec3[float32]

func TestAdd(t *testing.T) {
	a := TestVec{1, 2, 3}
	b := TestVec{8, 2, -4}

	result := Add(a, b)

	expected := TestVec{9, 4, -1}
	if !Equals(result, expected) {
		t.Errorf("Got %v vs expected %v", result, expected)
	}
}

func TestMul(t *testing.T) {
	a := TestVec{1, 2, 3}
	b := TestVec{8, 2, -4}

	result := Sub(a, b)

	expected := TestVec{-7, 0, 7}
	if !Equals(result, expected) {
		t.Errorf("Got %v vs expected %v", result, expected)
	}
}

func TestCross(t *testing.T) {
	a := TestVec{0, 0, 1}
	b := TestVec{1, 0, 0}

	result := Cross(a, b)
	expected := TestVec{0, 1, 0}

	if !Equals(result, expected) {
		t.Errorf("Got %v vs expected %v", result, expected)
	}
}

func TestDot(t *testing.T) {
	a := TestVec{1, 2, 3}
	b := TestVec{4, 5, 6}

	result := Dot(a, b)
	if result != 32 {
		t.Errorf("Got %v vs expected %v", result, 32)
	}
}

func TestLength(t *testing.T) {
	test := func(tested TestVec, expectedResult float32) {
		len := Length(tested)
		if len != expectedResult {
			t.Errorf("Got %f, expected %f", len, expectedResult)
		}
	}

	test(TestVec{1, 0, 0}, 1)
	test(TestVec{1, 1, 1}, float32(math.Sqrt(3)))
}

func TestNormalize(t *testing.T) {
	test := func(tested TestVec, expectedResult TestVec) {
		len := Normalize(tested)
		if len != expectedResult {
			t.Errorf("Got %v, expected %v", len, expectedResult)
		}
	}

	test(TestVec{1, 0, 0}, TestVec{1, 0, 0})
	test(TestVec{1, 1, 1}, TestVec{float32(1 / math.Sqrt(3)), float32(1 / math.Sqrt(3)), float32(1 / math.Sqrt(3))})
}

func TestMulScalar(t *testing.T) {
	test := func(tested TestVec, val float32, expectedResult TestVec) {
		result := MulScalar(tested, val)
		if result != expectedResult {
			t.Errorf("Got %v, expected %v", result, expectedResult)
		}
	}

	test(TestVec{1, 0, 0}, 3, TestVec{3, 0, 0})
	test(TestVec{2, 0, 0}, 3, TestVec{6, 0, 0})
	test(TestVec{1, 1, 1}, 3, TestVec{3, 3, 3})
}

func TestDivScalar(t *testing.T) {
	test := func(tested TestVec, val float32, expectedResult TestVec) {
		result := DivScalar(tested, val)
		if result != expectedResult {
			t.Errorf("Got %v, expected %v", result, expectedResult)
		}
	}

	test(TestVec{1, 0, 0}, 3, TestVec{1.0 / 3.0, 0, 0})
	test(TestVec{2, 0, 0}, 3, TestVec{2.0 / 3.0, 0, 0})
	test(TestVec{1, 1, 1}, 3, TestVec{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0})
}

func TestAddScalar(t *testing.T) {
	test := func(tested TestVec, val float32, expectedResult TestVec) {
		result := AddScalar(tested, val)
		if result != expectedResult {
			t.Errorf("Got %v, expected %v", result, expectedResult)
		}
	}

	test(TestVec{1, 0, 0}, 3, TestVec{4, 3, 3})
	test(TestVec{2, 0, 0}, 3, TestVec{5, 3, 3})
	test(TestVec{1, 1, 1}, 3, TestVec{4, 4, 4})
}

func TestSubScalar(t *testing.T) {
	test := func(tested TestVec, val float32, expectedResult TestVec) {
		result := SubScalar(tested, val)
		if result != expectedResult {
			t.Errorf("Got %v, expected %v", result, expectedResult)
		}
	}

	test(TestVec{1, 0, 0}, 3, TestVec{-2, -3, -3})
	test(TestVec{2, 0, 0}, 3, TestVec{-1, -3, -3})
	test(TestVec{1, 1, 1}, 3, TestVec{-2, -2, -2})
}
