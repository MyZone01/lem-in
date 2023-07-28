package lib

import (
	"lemin/models"
	"strconv"
	"strings"
)

func GetRoom(s string) (models.Room, bool) {
	ss := strings.Split(s, " ")
	room := models.Room{}
	if len(ss) == 3 {
		_, err1 := strconv.Atoi(ss[1])
		_, err2 := strconv.Atoi(ss[2])
		if err1 != nil || err2 != nil {
			return room, true
		}
		room.Name = ss[0]
		room.X = ss[1]
		room.Y = ss[2]
		return room, false
	}

	return room, true
}

func GetLink(s string) (models.Link, bool) {
	ss := strings.Split(s, "-")
	if len(ss) == 2 && ss[0] != ss[1] {
		return models.Link{
			From: ss[0],
			To:   ss[1],
		}, false
	}

	return models.Link{}, true
}

func SortPaths(Paths [][]string) {
	for i := 0; i < len(Paths)-1; {
		if len(Paths[i]) > len(Paths[i+1]) {
			tmp := Paths[i]
			Paths[i] = Paths[i+1]
			Paths[i+1] = tmp
			i = 0
		} else {
			i++
		}
	}
}

func StringToRoom(Paths [][]string, antFarm models.AntFarm) []models.Path {
	var PathsRoom []models.Path

	for _, path := range Paths {
		_rooms := []models.Room{}
		for _, room := range path {
			if room != antFarm.Start.Name {
				_rooms = append(_rooms, antFarm.Rooms[room])
			}
		}
		PathsRoom = append(PathsRoom, models.Path{Rooms: _rooms})
	}
	return PathsRoom
}
