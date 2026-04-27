package zones

type Terrain int

const (
	Water Terrain = iota
	Plain
	Tree
	Rock
	Rubble
	Wall
	Road
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
	Up
	Down
)

type Cell struct {
	TerrainType Terrain
	IsPassable  bool
	BlocksLOS   bool
}

type Zone struct {
	ID          int
	Name        string
	Description string
	XPos        int
	YPos        int
	Zmap        [][]Cell
	ZoneLog     string            // maybe should use like a fixed length array
	Neighbors   map[Direction]int // we need to talk about this because in my version you hit a zone line and then zip on to the next zone remember not a room in a normal mud
	Events      chan ZoneEvent    // players entering, actions, ticks
}
