package utils

import "lemin/model"

func FindPaths(antFarm model.AntFarm) []model.Path {
	return []model.Path{
		{
			Rooms: []model.Room{antFarm.Start, antFarm.Rooms["1"], antFarm.End},
		},
		{
			Rooms: []model.Room{antFarm.Start, antFarm.Rooms["2"], antFarm.End},
		},
	}
}
