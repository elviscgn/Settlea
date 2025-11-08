package orientation

import "math"

const (
	PointyLayout = "pointy"
	FlatLayout   = "flat"
)

type Orientation struct {
	ForwardMatrix [][]float64
	InvMatrix     [][]float64
	StartAngle    float64
}

func MakeOrientation(layoutType string) Orientation {
	switch layoutType {
	case PointyLayout:
		return Orientation{
			ForwardMatrix: [][]float64{
				{math.Sqrt(3.0), math.Sqrt(3.0) / 2.0},
				{0.0, 3.0 / 2.0},
			},
			InvMatrix: [][]float64{
				{math.Sqrt(3.0) / 3.0, -1.0 / 3.0},
				{0.0, 2.0 / 3.0},
			},
			StartAngle: 0.5,
		}
	case FlatLayout:
		return Orientation{
			ForwardMatrix: [][]float64{
				{3.0 / 2.0, 0.0},
				{math.Sqrt(3.0) / 2.0, math.Sqrt(3.0)},
			},
			InvMatrix: [][]float64{
				{2.0 / 3.0, 0.0},
				{-1.0 / 3.0, math.Sqrt(3.0) / 3.0},
			},
			StartAngle: 0.0,
		}
	default:
		panic("invalid layout type")
	}
}
