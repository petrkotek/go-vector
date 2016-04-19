package vec_test

import (
	"testing"

	"github.com/petrkotek/go-vector/2d/32bit/vec"
)

func TestNew(t *testing.T) {
	var x = float32(1.1)
	var y = float32(2.2)

	v := vec.New(x, y)
	if v.X != x || v.Y != y {
		t.Fail()
	}
}
