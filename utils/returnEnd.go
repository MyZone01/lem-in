package utils

import (
	"errors"
	"lemin/model"
	"strings"
)

func ReturnEnd(tab []string) (model.End, error) {
	response := model.End{}
	if len(tab) > 0 {
		response := model.End{}
		for i := 0; i < len(tab); i++ {
			if tab[i] == "##end" {
				if i+1 != len(tab) {
					response = MapEnd(tab[i+1])

					return response, nil
				} else {
					return response, errors.New("Invalid Syntax")
				}
			}
		}
	}
	return response, errors.New("Empty table")
}

func MapEnd(s string) model.End {
	ss := strings.Split(s, " ")

	tab := model.End {
		Name: ss[0],
		XEndRoom: ss[1],
		YEndRoom: ss[2],
	}

	return tab
}