package utils

import (
	"errors"
	"lemin/model"
	"strconv"
	"strings"
)

func GetRooms(tab []string, start model.Room, end model.Room) (map[string]model.Room, error) {
	tabFinal := map[string]model.Room{}
	if len(tab) > 0 {
		tabFinal := map[string]model.Room{}

		for _, val := range tab {
			if IsRoom(val) {
				if IsValidRoom(val) {
					if MapStart(val).Name != start.Name && MapEnd(val).Name != end.Name {
						room := MappingRooms(val)
						tabFinal[room.Name] = room
					}
				} else {
					return tabFinal, errors.New("ERROR: invalid data format, Room bad formatted")
				}
			}
		}
		return tabFinal, nil
	}
	return tabFinal, errors.New("ERROR: invalid data format, Room bad formatted")
}

func IsRoom(s string) bool {
	ss := strings.Split(s, " ")
	return len(ss) == 3
}

func IsValidRoom(s string) bool {
	ss := strings.Split(s, " ")
	if len(ss) == 3 {
		_, err1 := strconv.Atoi(ss[1])
		_, err2 := strconv.Atoi(ss[2])

		if err1 != nil || err2 != nil {
			return false
		}

		return true
	} else {
		return false
	}
}

func MappingRooms(s string) model.Room {
	ss := strings.Split(s, " ")

	data := model.Room{
		Name: ss[0],
		X:    ss[1],
		Y:    ss[2],
	}

	return data
}
