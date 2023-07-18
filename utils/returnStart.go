package utils

import (
	"errors"
	"lemin/model"
	"strings"
)

func ReturnStart(tab []string) (model.Start, error) {
	response := model.Start{}
	if len(tab) > 0 {
		response := model.Start{}
		for i := 0; i < len(tab); i++ {
			if tab[i] == "##start" {
				if i+1 != len(tab) {
					response = MapStart(tab[i+1])

					return response, nil
				} else {
					return response, errors.New("Invalid Syntax")
				}
			}
		}
	}
	return response, errors.New("Empty table")
}

func MapStart(s string) model.Start {
	ss := strings.Split(s, " ")

	tab := model.Start {
		Name: ss[0],
		XStartRoom: ss[1],
		YStartRoom: ss[2],
	}

	return tab
}