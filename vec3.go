package vec3

import (
	"math"

	"golang.org/x/exp/constraints"
)

const EPS = 0.000001

type Flt interface {
	constraints.Float
}

func abs[Num Flt](val Num) Num {
	if val < 0 {
		return -val
	}
	return val
}

type Vec3[Num Flt] struct {
	X, Y, Z Num
}

type elementOp[Num Flt] func(Num, Num) Num

func piecewiseOp[Num Flt](v1 Vec3[Num], v2 Vec3[Num], op elementOp[Num]) Vec3[Num] {
	return Vec3[Num]{
		op(v1.X, v2.X),
		op(v1.Y, v2.Y),
		op(v1.Z, v2.Z),
	}
}

func constwiseOp[Num Flt](v1 Vec3[Num], val Num, op elementOp[Num]) Vec3[Num] {
	return Vec3[Num]{
		op(v1.X, val),
		op(v1.Y, val),
		op(v1.Z, val),
	}
}

func Dot[Num Flt](self Vec3[Num], other Vec3[Num]) Num {
	return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func Cross[Num Flt](self Vec3[Num], other Vec3[Num]) Vec3[Num] {
	return Vec3[Num]{
		X: self.Y*other.Z - other.Z*other.Y,
		Y: self.Z*other.X - other.X*other.Z,
		Z: self.X*other.Y - other.Y*other.X,
	}
}

func MulScal[Num Flt](self *Vec3[Num], val Num) Vec3[Num] {
	return Vec3[Num]{
		X: self.X * val,
		Y: self.Y * val,
		Z: self.Z * val,
	}
}

func Add[Num Flt](self Vec3[Num], other Vec3[Num]) Vec3[Num] {
	return piecewiseOp(self, other, func(a, b Num) Num { return a + b })
}

func Sub[Num Flt](self Vec3[Num], other Vec3[Num]) Vec3[Num] {
	return piecewiseOp(self, other, func(a, b Num) Num { return a - b })
}

func Length[Num Flt](self Vec3[Num]) Num {
	squaredSum := self.X * self.X
	squaredSum += self.Y * self.Y
	squaredSum += self.Z * self.Z
	return Num(math.Sqrt(float64(squaredSum)))
}

func Normalize[Num Flt](self Vec3[Num]) Vec3[Num] {
	len := Length(self)
	return Vec3[Num]{self.X / len, self.Y / len, self.Z / len}
}

func MulScalar[Num Flt](self Vec3[Num], val Num) Vec3[Num] {
	return constwiseOp(self, val, func(a, b Num) Num { return a * b })
}

func DivScalar[Num Flt](self Vec3[Num], val Num) Vec3[Num] {
	return constwiseOp(self, val, func(a, b Num) Num { return a / b })
}

func Equals[Num Flt](self Vec3[Num], other Vec3[Num]) bool {
	return abs(self.X-other.X) < EPS &&
		abs(self.Y-other.Y) < EPS &&
		abs(self.Z-other.Z) < EPS
}
