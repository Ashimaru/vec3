package vec3

import (
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

func Dot[Num Flt](self *Vec3[Num], other *Vec3[Num]) Num {
	return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func Cross[Num Flt](self *Vec3[Num], other *Vec3[Num]) Vec3[Num] {
	return Vec3[Num]{
		X: self.Y*other.Z - other.Z*other.Y,
		Y: self.Z*other.X - other.X*other.Z,
		Z: self.X*other.Y - other.Y*other.X,
	}
}

func Add[Num Flt](self *Vec3[Num], other *Vec3[Num]) Vec3[Num] {
	return Vec3[Num]{
		X: self.X + other.X,
		Y: self.Y + other.Y,
		Z: self.Z + other.Z,
	}
}

func Sub[Num Flt](self *Vec3[Num], other *Vec3[Num]) Vec3[Num] {
	return Vec3[Num]{
		X: self.X - other.X,
		Y: self.Y - other.Y,
		Z: self.Z - other.Z,
	}
}

func MulScal[Num Flt](self *Vec3[Num], val Num) Vec3[Num] {
	return Vec3[Num]{
		X: self.X * val,
		Y: self.Y * val,
		Z: self.Z * val,
	}
}

func Equals[Num Flt](self *Vec3[Num], other *Vec3[Num]) bool {
	return abs((self.X-other.X)) < EPS &&
		abs((self.Y-other.Y)) < EPS &&
		abs((self.Z-other.Z)) < EPS
}
