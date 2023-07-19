package utils

import "lemin/model"

func StringToRoom(Paths [][]string, antFarm model.AntFarm) []model.Path {
	var PathsRoom []model.Path

	for _, path := range Paths {
		_rooms := []model.Room{}
		for _, room := range path {
			if room != antFarm.Start.Name {
				_rooms = append(_rooms, antFarm.Rooms[room])
			}
		}
		PathsRoom = append(PathsRoom, model.Path{Rooms: _rooms})
	}
	return PathsRoom
}