package grid

import (
	"settlea/pkg/bestagons/edge"
	"settlea/pkg/bestagons/hex"
	"settlea/pkg/bestagons/orientation"
	"settlea/pkg/bestagons/screen"
	"settlea/pkg/bestagons/vertex"
	"testing"
)

func TestVertexToPixel(t *testing.T) {
	layout_pointy := orientation.MakeOrientation(orientation.PointyLayout)

	layout := Layout{layout_pointy, screen.MakeScreenCoord(50, 50), screen.MakeScreenCoord(0, 0)}
	vertexTest := vertex.Vertex{Q: 0, R: 0, Direction: "N"}
	pixelCoord := layout.VertexToPixel(vertexTest)

	expected := screen.MakeScreenCoord(28.86751345948129, -8.333333333333332)

	if pixelCoord != expected {
		t.Errorf("Expected %v, got %v", expected, pixelCoord)
	}

}

func TestHexNeighbours(t *testing.T) {
	layout_pointy := orientation.MakeOrientation(orientation.PointyLayout)

	layout := Layout{layout_pointy, screen.MakeScreenCoord(50, 50), screen.MakeScreenCoord(0, 0)}
	hexTest := hex.Hex{Q: 0, R: 0}
	neighbours := layout.neighbours(hexTest)

	expected := []hex.Hex{
		*hex.NewHex(0, -1),
		*hex.NewHex(1, -1),
		*hex.NewHex(1, 0),
		*hex.NewHex(0, 1),
		*hex.NewHex(-1, 1),
		*hex.NewHex(-1, 0),
	}

	if len(neighbours) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, neighbours)
	}

	for i, n := range neighbours {
		if n != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], n)
		}
	}
}

func TestVertices(t *testing.T) {
	layout_pointy := orientation.MakeOrientation(orientation.PointyLayout)

	layout := Layout{layout_pointy, screen.MakeScreenCoord(50, 50), screen.MakeScreenCoord(0, 0)}
	hexTest := hex.Hex{Q: 1, R: 1}

	vertices := layout.Vertices(hexTest)

	expected := []vertex.Vertex{
		*vertex.NewVertex(1, 1, vertex.North),
		*vertex.NewVertex(2, 0, vertex.South),
		*vertex.NewVertex(1, 2, vertex.North),
		*vertex.NewVertex(1, 1, vertex.South),
		*vertex.NewVertex(0, 2, vertex.North),
		*vertex.NewVertex(1, 0, vertex.South),
	}

	if len(vertices) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, vertices)
	}

	for i, v := range vertices {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], v)
		}
	}
}

func TestEdges(t *testing.T) {
	layout_pointy := orientation.MakeOrientation(orientation.PointyLayout)
	layout := Layout{layout_pointy, screen.MakeScreenCoord(50, 50), screen.MakeScreenCoord(0, 0)}

	hexTest := hex.Hex{Q: 1, R: 1}

	edges := layout.Edges(hexTest)

	expected := []edge.Edge{
		*edge.NewEdge(1, 1, edge.NorthWest),
		*edge.NewEdge(1, 1, edge.NorthEast),
		*edge.NewEdge(2, 1, edge.West),
		*edge.NewEdge(1, 2, edge.NorthWest),
		*edge.NewEdge(0, 2, edge.NorthEast),
		*edge.NewEdge(1, 1, edge.West),
	}

	if len(edges) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, edges)
	}

	for i, e := range edges {
		if e != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], e)
		}
	}

}
