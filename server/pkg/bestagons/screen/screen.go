// adding this JIC, i might let the client handle this instead;
package screen

type ScreenCoord struct {
	X float64
	Y float64
}

// func NewScreenCoord(x int, y int) *ScreenCoord {
// 	return &ScreenCoord{x, y}
// }

func MakeScreenCoord(x float64, y float64) ScreenCoord {
	return ScreenCoord{x, y}
}

// func (s ScreenCoord) equals(other ScreenCoord) bool {
// 	return s.X == other.X && s.Y == other.Y
// }
