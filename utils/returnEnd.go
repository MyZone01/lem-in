package utils

import (
	"errors"
	"fmt"
	"lemin/model"
	"strings"
)

func ReturnEnd(tab []string) (model.Room, error) {
	response := model.Room{}
	if len(tab) > 0 {
		response := model.Room{}
		for i := 0; i < len(tab); i++ {
			if tab[i] == "##end" {
				if i+1 != len(tab) {
					response = MapEnd(tab[i+1])
					if len(response.Name) == 0 || len(response.X) == 0  || len(response.Y) == 0{
						return response, errors.New("ERROR: invalid data format, no end room found")
					}
					return response, nil
				} else {
					return response, errors.New("ERROR: invalid data format, no end room found")
				}
			}
		}
	}
	return response, errors.New("ERROR: invalid data format, no end room found")
}

func MapEnd(s string) model.Room {
	tab := model.Room{}
	if len(s) > 0 {
		ss := strings.Fields(s)
		fmt.Println(ss)

		if len(ss) == 3 {
			tab := model.Room{
				Name: ss[0],
				X:    ss[1],
				Y:    ss[2],
			}
		
			return tab
		}
	}
	return tab
}
