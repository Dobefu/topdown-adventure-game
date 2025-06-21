package vectors

import (
	"math"
)

type IVector3 interface {
	Add(vec Vector3)
	Sub(vec Vector3)
	Mul(vec Vector3)
	Div(vec Vector3)
	Normalize() (vec Vector3)
	ToVector2() Vector2
}

type Vector3 struct {
	IVector3

	X float64
	Y float64
	Z float64
}

func (v *Vector3) Add(vec Vector3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

func (v *Vector3) Sub(vec Vector3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

func (v *Vector3) Mul(vec Vector3) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
}

func (v *Vector3) Div(vec Vector3) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
}

func (v *Vector3) Normalize() (vec Vector3) {
	vec = *v
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

	if magnitude != 0 {
		vec.X /= magnitude
		vec.Y /= magnitude
		vec.Z /= magnitude
	}

	return vec
}

func (v Vector3) ToVector2() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}
