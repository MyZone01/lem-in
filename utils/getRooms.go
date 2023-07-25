package utils

import (
	"errors"
	"lemin/model"
	"strconv"
	"strings"
)

func GetRooms(tab []string, start model.Room, end model.Room) (map[string]model.Room, error) {
	tabFinal := map[string]model.Room{}
	var tabRoom []string
	if len(tab) > 0 {
		tabFinal := map[string]model.Room{}

		for _, val := range tab {
			if IsRoom(val) {
				if IsValidRoom(val) {
					if MapStart(val).Name != start.Name && MapEnd(val).Name != end.Name {
						room := MappingRooms(val)
						tabFinal[room.Name] = room
						tabRoom = append(tabRoom, room.Name)
					}
				} else {
					return tabFinal, errors.New("ERROR: invalid data format, Room bad formatted")
				}
			}
		}
		if AllRoomUnique(tabRoom) {
			return tabFinal, nil
		} else {
			return tabFinal , errors.New("ERROR: Two Room can't have the same name")
		}
	}
	return tabFinal, errors.New("ERROR: invalid data format, Room bad formatted")
}

func IsRoom(s string) bool {
	ss := strings.Fields(s)
	return len(ss) == 3
}

func AllRoomUnique(tab []string) bool {
	for i,_ := range tab {
		for j, _ := range tab {
			if tab[i] == tab[j] && i != j {
				return false
			}
		}
	}
	return true
}

func IsValidRoom(s string) bool {
	ss := strings.Fields(s)
	if len(ss) == 3 {
		_, err1 := strconv.Atoi(ss[1])
		_, err2 := strconv.Atoi(ss[2])


		if err1 != nil || err2 != nil {
			return false
		}

		if strings.EqualFold(string(ss[0][0]), "L") ||  strings.EqualFold(string(ss[0][0]), "#") {
			return false
		}

		return true
	} else {
		return false
	}
}

func MappingRooms(s string) model.Room {
	ss := strings.Fields(s)

	data := model.Room{
		Name: ss[0],
		X:    ss[1],
		Y:    ss[2],
	}

	return data
}
