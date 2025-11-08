package hex

type Hex struct {
	Q int
	R int
	S int
}

func NewHex(q int, r int) *Hex {
	return &Hex{q, r, -q - r}
}

func MakeHex(q int, r int) Hex {
	return Hex{q, r, -q - r}
}

func (h Hex) GetNeighbour(direction int) Hex {
	switch direction {
	case 0:
		return MakeHex(h.Q, h.R-1)
	case 1:
		return MakeHex(h.Q+1, h.R-1)
	case 2:
		return MakeHex(h.Q+1, h.R)
	case 3:
		return MakeHex(h.Q, h.R+1)
	case 4:
		return MakeHex(h.Q-1, h.R+1)
	case 5:
		return MakeHex(h.Q-1, h.R)
	default:
		panic("invalid direction")
	}
}
