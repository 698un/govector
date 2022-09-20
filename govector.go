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

func (v1 *MVector) Set(inpX, inpY, inpZ float32) {
	v1.X = inpX
	v1.Y = inpY
	v1.Z = inpZ
}

func Sum2(v1, v2 MVector) MVector {
	var res MVector
	res.X = v1.X + v2.X
	res.Y = v1.Y + v2.Y
	res.Z = v1.Z + v2.Z
	return res

}

func Sum(v1 []MVector) MVector {

	res := NewVector()

	for i := 0; i < v1.length; i++ {
		res.X += v1[i].X
		res.Y += v1[i].Y
		res.Z += v1[i].Z
	} //next i

	return res
}

func MultScalar(v1, v2 MVector) float32 {

	return 0 +
		v1.X*v2.X +
		v1.Y*v2.Y +
		v1.Z*v2.Z

} //MultVector

func MultVector(v1, v2 MVector) MVector {

	var res MVector

	res.X = v1.Y*v2.Z - v1.Z*v2.Y
	res.Y = v1.Z*v2.X - v1.X*v2.Z
	res.Z = v1.X*v2.Y - v1.Y*v2.X

	return res
} //MultVector
