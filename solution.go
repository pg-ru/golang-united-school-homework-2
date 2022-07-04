package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesNum int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum sidesNum) float64 {
	var square float64

	switch sidesNum {
	case SidesTriangle:
		p := (SidesTriangle * sideLen) / 2
		square = math.Sqrt(p * (math.Pow(p-sideLen, 3)))
	case SidesSquare:
		square = math.Pow(sideLen, 2)
	case SidesCircle:
		square = math.Pi * math.Pow(sideLen, 2)
	default:
		square = 0
	}

	return square
}
