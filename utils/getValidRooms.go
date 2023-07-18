package utils

import (
	"errors"
	"lemin/model"
)

func GetValidRooms(Rooms map[string]model.Room, Links []model.Link) (map[string]model.Room, error) {
	RoomsFinal := map[string]model.Room{}
	for _, room := range Rooms {
		if GoodRoom(room, Links) {
			RoomsFinal[room.Name] = room
		}
	}
	if len(Rooms) > 0 {
		return RoomsFinal, nil
	}
	return RoomsFinal, errors.New("no valid room")
}

func GoodRoom(Room model.Room, Links []model.Link) bool {
	for _, val := range Links {
		if val.Start == Room.Name || val.End == Room.Name {
			return true
		}
	}
	return false
}
