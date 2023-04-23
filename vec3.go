package mmich

import (
	"golang.org/x/exp/constraints"
)

type Flt interface {
	constraints.Float
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
