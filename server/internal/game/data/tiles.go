package data

import (
	"settlea/pkg/bestagons/edge"
	"settlea/pkg/bestagons/grid"
	"settlea/pkg/bestagons/hex"
	"settlea/pkg/bestagons/screen"
	"settlea/pkg/bestagons/vertex"
	"settlea/pkg/utils"
	"strconv"
	"time"
)

type Tile struct {
	hex.Hex
	Type    string
	Token   int
	Blocked bool
	Coords  screen.ScreenCoord
}

func (t *Tile) setToken(token int) {
	t.Token = token
}

func (t Tile) SetType(tileType string) Tile {
	t.Type = tileType
	return t
}

func (t Tile) SetBlocked(blocked bool) Tile {
	t.Blocked = blocked
	return t
}

type Board struct {
	Tiles utils.Set[Tile]
}

func NewBoard(N int) []*Tile {
	settleaMap := GenerateHexagonMap(N)

	return settleaMap
}

func GenerateHexagonMap(N int) []*Tile {
	tiles := make([]*Tile, 0)

	if N != 2 {
		panic("Not implemented")
	}

	resources := []string{}
	resources = append(resources, utils.Repeat("wood", 4)...)
	resources = append(resources, utils.Repeat("sheep", 4)...)
	resources = append(resources, utils.Repeat("wheat", 4)...)
	resources = append(resources, utils.Repeat("brick", 3)...)
	resources = append(resources, utils.Repeat("ore", 3)...)
	resources = append(resources, "desert")
	shuffledResources := utils.Shuffle(resources)

	for q := -N; q <= N; q++ {
		r1 := max(-N, -q-N)
		r2 := min(N, -q+N)

		for r := r1; r <= r2; r++ {
			// Popping the resource using slice manipulation
			resource := shuffledResources[len(shuffledResources)-1]          // Get the last resource
			shuffledResources = shuffledResources[:len(shuffledResources)-1] // Remove the last resource

			tile := &Tile{
				Hex:     hex.MakeHex(q, r),
				Type:    resource,
				Token:   0,
				Blocked: false,
			}
			tiles = append(tiles, tile)
		}
	}
	return tiles
}

func StartValidation(tiles []*Tile) ([]*Tile, int, time.Duration) {
	legal := false
	iteration := 0
	startTime := time.Now()

	for !legal && iteration < 1000 {
		iteration++
		tokens := []string{"2"}
		tokens = append(tokens, utils.Repeat("3", 2)...)
		tokens = append(tokens, utils.Repeat("4", 2)...)
		tokens = append(tokens, utils.Repeat("5", 2)...)
		tokens = append(tokens, utils.Repeat("6", 2)...)
		tokens = append(tokens, utils.Repeat("8", 2)...)
		tokens = append(tokens, utils.Repeat("9", 2)...)
		tokens = append(tokens, utils.Repeat("10", 2)...)
		tokens = append(tokens, utils.Repeat("11", 2)...)
		tokens = append(tokens, "12")

		nonDesertCount := 0
		for _, tile := range tiles {
			if tile.Type != "desert" {
				nonDesertCount++
			}
		}

		if nonDesertCount != len(tokens) {
			panic("Mismatch between number of tokens and non-desert tiles, got " + strconv.Itoa(nonDesertCount) + " and " + strconv.Itoa(len(tokens)))
		}

		shuffledTokens := utils.Shuffle(tokens)
		for _, tile := range tiles {
			if tile.Type == "desert" {
				continue
			}

			// Using the simplified popping method
			token := shuffledTokens[len(shuffledTokens)-1]          // Get the last token
			shuffledTokens = shuffledTokens[:len(shuffledTokens)-1] // Remove the last token

			tokenInt, err := strconv.Atoi(token)
			if err != nil {
				panic("Invalid token conversion: " + err.Error())
			}
			tile.setToken(tokenInt)
		}

		legal = validateTiles(tiles)
	}

	duration := time.Since(startTime)
	return tiles, iteration, duration
}

func validateTiles(hex_map []*Tile) bool {
	// Convert set to a map for faster lookups
	hex_dict := make(map[hex.Hex]*Tile)
	for _, tile := range hex_map {
		hex_dict[tile.Hex] = tile
	}

	for _, tile := range hex_map {
		for i := 0; i < 6; i++ {
			neighbour := tile.Hex.GetNeighbour(i)
			neighbor_tile, exists := hex_dict[neighbour]

			if exists {
				// Check if neighboring tokens are identical
				if tile.Token == neighbor_tile.Token {
					return false
				}
				// Check if the adjacent tokens are 6 and 8, which are not allowed to be neighbors
				if (tile.Token == 6 && neighbor_tile.Token == 8) || (tile.Token == 8 && neighbor_tile.Token == 6) {
					return false
				}
			}
		}
	}

	return true
}

func GenerateVertices(layout grid.Layout, tiles []*Tile) []*vertex.Vertex {
	uniqueVertices := utils.Set[vertex.Vertex]{}

	for _, tile := range tiles {
		hex := tile.Hex
		tileVertices := layout.Vertices(hex)
		for _, v := range tileVertices {
			uniqueVertices.Add(v)
		}
	}

	vertices := make([]*vertex.Vertex, 0, uniqueVertices.Size())
	for v := range uniqueVertices {
		vertices = append(vertices, &v)
	}

	return vertices
}

func GenerateEdges(layout grid.Layout, tiles []*Tile) []*edge.Edge {
	uniqueEdges := utils.Set[edge.Edge]{}

	for _, tile := range tiles {
		hex := tile.Hex
		tileEdges := layout.Edges(hex)
		for _, e := range tileEdges {
			uniqueEdges.Add(e)

		}
	}

	edges := make([]*edge.Edge, 0, uniqueEdges.Size())
	for e := range uniqueEdges {
		edges = append(edges, &e)

	}

	return edges
}

func GenerateCorners(layout grid.Layout, vertices []*vertex.Vertex) []*Corner {
	corners := make([]*Corner, 0)

	for _, v := range vertices {
		corner := &Corner{
			Vertex:    *v,
			Structure: nil,
			IsPort:    false,
			PortType:  "",
		}
		corners = append(corners, corner)
	}

	return corners
}
