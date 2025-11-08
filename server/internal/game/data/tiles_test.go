package data

import (
	"settlea/pkg/bestagons/grid"
	"settlea/pkg/bestagons/orientation"
	"settlea/pkg/bestagons/screen"
	"testing"
	// "time"
)

func TestGenerateHexagonalMap(t *testing.T) {

	hex_map := GenerateHexagonMap(2)

	map_size := len(hex_map)
	expected_size := 19

	if map_size != expected_size {
		t.Errorf("Expected %v, got %v", expected_size, map_size)
	}
}

func TestValidateTiles(t *testing.T) {

	tiles := GenerateHexagonMap(2)

	// Run StartValidation to assign tokens and get a valid configuration
	result, _, _ := StartValidation(tiles)

	// Validate the result using validateTiles
	legal := validateTiles(result)

	if !legal {
		t.Errorf("Expected tiles to be valid but found conflicts")
	}
}

func TestGenerateVertices(t *testing.T) {

	layout := grid.Layout{
		Orientation: orientation.MakeOrientation(orientation.PointyLayout), // kinda redundant since all layouts will b pointy
		Origin:      screen.MakeScreenCoord(0, 0),
		Size:        screen.MakeScreenCoord(92, 92),
	}
	tiles := GenerateHexagonMap(2)
	vertices := GenerateVertices(layout, tiles)

	expected_len := 54

	if len(vertices) != expected_len {
		t.Errorf("Expected %v, got %v", expected_len, len(vertices))
	}

}
