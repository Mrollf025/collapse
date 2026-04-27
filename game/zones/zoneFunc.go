package zones

func BuildTestZone() Zone {
	baseCell := Cell{
		TerrainType: 1,
		IsPassable:  true,
		BlocksLOS:   false,
	}

	zmap := make([][]Cell, 10)

	for i := 0; i < 10; i++ {
		row := make([]Cell, 10)
		for j := 0; j < 0; j++ {
			row[j] = baseCell
		}
		zmap[i] = row
	}

	return Zone{
		ID:          1,
		Name:        "Test ZONE",
		Description: "This is a hard set zone for testing functionality",
		XPos:        0,
		YPos:        0,
		Zmap:        zmap,
		ZoneLog:     "",
		Neighbors:   map[Direction]int{},
	}
}
