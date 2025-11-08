package vertex

type Vertex struct {
	Q         int
	R         int
	Direction string
}

const (
	North = "N"
	South = "S"
)

func NewVertex(q, r int, dir string) *Vertex {
	return &Vertex{Q: q, R: r, Direction: dir}
}

func (v Vertex) IsValid() bool {
	return v.Direction == North || v.Direction == South
}
