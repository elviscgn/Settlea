package edge

type Edge struct {
	Q         int
	R         int
	Direction string
}

const (
	NorthEast = "NE"
	NorthWest = "NW"
	West      = "W"
)

func NewEdge(q, r int, dir string) *Edge {
	return &Edge{Q: q, R: r, Direction: dir}
}

func (e Edge) IsValid() bool {
	return e.Direction == NorthEast || e.Direction == NorthWest || e.Direction == West
}
