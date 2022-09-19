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

func (v1 *MVector) SetValues(inpX, inpY, inpZ float32) {
	v1.X = inpX
	v1.Y = inpY
	v1.Z = inpZ
}

func SumTwo(v1, v2 MVector) MVector {
	var res MVector

	res.X = v1.X + v2.X
	res.Y = v1.Y + v2.Y
	res.Z = v1.Z + v2.Z
	return res

}

func MultScalar(v1, v2 MVector) float32 {

	return 0 +
		v1.X*v2.X +
		v1.Y*v2.Y +
		v1.Z*v2.Z

} //MultVector
