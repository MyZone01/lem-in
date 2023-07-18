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
				if IsValide(tab[i]) {
					tabFinal = append(tabFinal, Mapping(tab[i]))
				} else {
					return nil, errors.New("Invalid syntax, Bad Format Tunnel")
				}
			}
		}
		return tabFinal, nil
	}
	return nil, errors.New("Empty Table")
}

func IsValide(s string) bool {
	ss := strings.Split(s, "-")

	if ss[0] != ss[1] {
		return true
	}
	return false
}

func IsLink(s string) bool {
	ss := strings.Split(s, "-")

	if len(ss) == 2 {
		return true
	}
	return false
}

func Mapping(s string) model.Link {
	ss := strings.Split(s, "-")

	data := model.Link {
		X_link: ss[0],
		Y_link: ss[1],
	}

	return data
}