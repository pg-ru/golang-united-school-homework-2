package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesNum int

const TRIANGLE = 3
const SQUARE = 4
const CIRCLE = 0

func CalcSquare(sideLen float64, sidesNum sidesNum) float64 {
	var square float64

	switch sidesNum {
	case TRIANGLE:
		square = TRIANGLE * sideLen
	case SQUARE:
		square = SQUARE * sideLen
	case CIRCLE:
		square = math.Pi * math.Pow(sideLen, 2)
	default:
		square = 0
	}

	return square
}
