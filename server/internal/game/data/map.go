package data

import (
	"settlea/pkg/bestagons/edge"
	"settlea/pkg/bestagons/grid"
	"settlea/pkg/bestagons/orientation"
	"settlea/pkg/bestagons/screen"
)

var base_layout = grid.Layout{
	Orientation: orientation.MakeOrientation(orientation.PointyLayout), // kinda redundant since all layouts will b pointy
	Origin:      screen.MakeScreenCoord(0, 0),
	Size:        screen.MakeScreenCoord(92, 92),
}

type SettleaMap struct {
	Layout  grid.Layout
	Tiles   []*Tile
	Corners []*Corner
	Edges   []*edge.Edge
}

func (s *SettleaMap) NewMap(style string) *SettleaMap {
	switch style {
	case "base":
		return s.setupMap(base_layout)
	// more map styles to come
	default:
		panic("not implemented")
	}

}

func (s *SettleaMap) setupMap(layout grid.Layout) *SettleaMap {
	tiles := GenerateHexagonMap(2)
	vertices := GenerateVertices(layout, tiles)
	edges := GenerateEdges(layout, tiles)

	corners := GenerateCorners(layout, vertices)

	// Run StartValidation to assign tokens and get a valid configuration
	validTiles, _, _ := StartValidation(tiles)

	return &SettleaMap{
		layout,
		validTiles,
		corners,
		edges,
	}

}
