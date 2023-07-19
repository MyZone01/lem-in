package utils

import (
	"errors"
	"lemin/model"
	"strings"
)

func ReturnStart(tab []string) (model.Room, error) {
	response := model.Room{}
	if len(tab) > 0 {
		response := model.Room{}
		for i := 0; i < len(tab); i++ {
			if tab[i] == "##start" {
				if i+1 != len(tab) {
					response = MapStart(tab[i+1])

					return response, nil
				} else {
					return response, errors.New("invalid syntax")
				}
			}
		}
	}
	return response, errors.New("empty table")
}

func MapStart(s string) model.Room {
	ss := strings.Split(s, " ")

	tab := model.Room{
		Name: ss[0],
		X:    ss[1],
		Y:    ss[2],
	}

	return tab
}
