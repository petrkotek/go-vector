package vec

import (
	"fmt"
)

type Vector2D struct {
	X float32
	Y float32
}

func New(x, y float32) Vector2D {
	return Vector2D{x, y}
}

func FromScalar(v float32) Vector2D {
	return Vector2D{v, v}
}

func Zero() Vector2D {
	return Vector2D{0, 0}
}

func Unit() Vector2D {
	return Vector2D{1, 1}
}

func (v Vector2D) Copy() Vector2D {
	return Vector2D{v.X, v.Y}
}

func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{v.X * v2.X, v.Y * v2.Y}
}

func (v Vector2D) Divide(v2 Vector2D) Vector2D {
	return Vector2D{v.X / v2.X, v.Y / v2.Y}
}

func (v Vector2D) MultiplyScalar(s float32) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
}

func (v Vector2D) DivideScalar(s float32) Vector2D {
	return Vector2D{v.X / s, v.Y / s}
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%v:%v", v.X, v.Y)
}
