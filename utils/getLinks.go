package utils

import (
	"errors"
	"lemin/model"
	"strings"
)

func GetLink(tab []string) ([]model.Link, error) {
	if len(tab) > 0 {
		var tabFinal []model.Link
		
		for i := 0; i < len(tab); i++ {
			if IsLink(tab[i]) {
				if IsValid(tab[i]) {
					tabFinal = append(tabFinal, Mapping(tab[i]))
					
				} else {
					return nil, errors.New("ERROR: invalid syntax, bad format tunnel")
				}
			} else if !IsRoom(tab[i]) && tab[i] != "##start" && tab[i] != "##end" && string(tab[i][0]) != "#" {
					return nil, errors.New("ERROR: invalid syntax, bad format")
			}
		}
		if CheckLinkIsUnique(tabFinal) {
			return tabFinal, nil
		} else {
			return tabFinal, errors.New("ERROR: Link must be unique")
		}
	}
	return nil, errors.New("ERROR: empty table")
}

func IsValid(s string) bool {
	ss := strings.Split(s, "-")
	if len(ss) == 2 {
		return ss[0] != ss[1] 
	}
	return false
}

func IsLink(s string) bool {
	ss := strings.Split(s, "-")
	return len(ss) == 2 
}

func CheckLinkIsUnique(tab []model.Link) bool {
	for i, _ := range tab {
		for j, _ := range tab {
			if ((tab[i].From == tab[j].From && tab[i].To == tab[j].To) || (tab[i].To == tab[j].From && tab[i].From == tab[j].To)) && i != j {
				return false
			}
		}
	}
	return true
}

func Mapping(s string) model.Link {
	ss := strings.Split(s, "-")

	ss[0] = strings.Trim(ss[0], " ")
	ss[1] = strings.Trim(ss[1], " ")

	data := model.Link {
		From: ss[0],
		To: ss[1],
	}

	return data
}
