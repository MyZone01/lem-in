package utils

import (
	"errors"
	"lemin/model"
)

func GetValidRooms(Rooms []model.Room, Links []model.Link) ([]model.Room, error) {
	RoomsFinal := []model.Room{}
	for i:=0; i < len(Rooms) ; i++ {
		if GoodRoom(Rooms[i], Links) {
			RoomsFinal = append(RoomsFinal, Rooms[i])
		}
	}
	if len(Rooms) > 0 {
		return RoomsFinal, nil
	}
	return RoomsFinal, errors.New("No Valid Room")
}

func GoodRoom(Room model.Room, Links []model.Link) bool {
	for _, val := range Links {
		if val.From == Room.Name || val.To == Room.Name {
			return true
		}
	}
	return false
}
