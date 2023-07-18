package utils

import (
	"errors"
	"lemin/model"
	"strconv"
	"strings"
)

func GetRooms(tab []string) (map[string]model.Room, error) {
	tabFinal := map[string]model.Room{}
	if len(tab) > 0 {
		StartRoom, _ := ReturnStart(tab)
		EndRoom, _ := ReturnEnd(tab)

		tabFinal := map[string]model.Room{}

		for _, val := range tab {
			if IsRoom(val) {

				if IsValidRoom(val) {

					if MapStart(val) != StartRoom && MapEnd(val) != EndRoom {
						room := MappingRooms(val)
						tabFinal[room.Name] = room
					}
				} else {
					return tabFinal, errors.New("error syntax, room syntax wrong")
				}
			}
		}
		return tabFinal, nil
	}
	return tabFinal, errors.New("empty file")
}

func IsRoom(s string) bool {
	ss := strings.Split(s, " ")
	return len(ss) == 3
}

func IsValidRoom(s string) bool {
	ss := strings.Split(s, " ")

	_, err1 := strconv.Atoi(ss[1])
	_, err2 := strconv.Atoi(ss[2])

	if err1 != nil || err2 != nil {
		return false
	}

	return true
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
