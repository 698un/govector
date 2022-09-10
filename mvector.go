/*
This modul is vector unit for GO

*/

package govector

type MVector struct {
	X float32
	Y float32
	Z float32
}

func NewVector() MVector {
	var res MVector
	res.X = 0
	res.Y = 0
	res.Z = 0
	return res

}

func SumTest(x, y int) int {
	return x + y
}
