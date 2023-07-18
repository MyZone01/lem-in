package utils

import (
	"errors"
	"lemin/model"
	"strconv"
	"strings"
)

func GetRooms(tab []string) ([]model.Room, error) {
	tabFinal := []model.Room{}
	if len(tab) > 0 {
		StartRoom, _ := ReturnStart(tab)
		EndRoom, _ := ReturnEnd(tab)

		tabFinal := []model.Room{}

		for _, val := range tab {
			if IsRoom(val) {
				
				if IsValideRoom(val) {
					
					if MapStart(val) != StartRoom && MapEnd(val) != EndRoom {
						
						tabFinal = append(tabFinal, MappingRooms(val))
					}
				} else {
					return tabFinal, errors.New("Error syntax, Room syntax wrong")
				}
			}
		}
		return tabFinal, nil
	}
	return tabFinal, errors.New("Empty File")
}

func IsRoom(s string) bool {
	ss := strings.Split(s, " ")

	if len(ss) == 3 {
		return true
	}
	return false
}

func IsValideRoom(s string) bool {
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

	data := model.Room {
		Name: ss[0],
		X_room: ss[1],
		Y_room : ss[2],
	}

	return data
}