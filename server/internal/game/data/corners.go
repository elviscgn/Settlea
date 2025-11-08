package data

import (
	"settlea/pkg/bestagons/screen"
	"settlea/pkg/bestagons/vertex"
)

type Corner struct {
	vertex.Vertex
	Coords screen.ScreenCoord
	// Coords    screen.ScreenCoord // currently works better if the client computes it based on screen width and height
	Structure *structure
	IsPort    bool
	PortType  string // 3:1, 2:1->wood,brick,sheep,ore,wheat
}

func (c *Corner) SetStructure(s *structure) {
	c.Structure = s
}

func (c *Corner) RemoveStructure() {
	c.Structure = nil
}

func (c *Corner) SetPort(isPort bool, portType string) {
	c.IsPort = isPort
	c.PortType = portType
}

func (c *Corner) HasStructure() bool {
	return c.Structure != nil
}
