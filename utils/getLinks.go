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
					return nil, errors.New("invalid syntax, bad format tunnel")
				}
			}
		}
		return tabFinal, nil
	}
	return nil, errors.New("empty table")
}

func IsValid(s string) bool {
	ss := strings.Split(s, "-")
	return ss[0] != ss[1]
}

func IsLink(s string) bool {
	ss := strings.Split(s, "-")
	return len(ss) == 2
}

func Mapping(s string) model.Link {
	ss := strings.Split(s, "-")

	data := model.Link{
		Start: ss[0],
		End:   ss[1],
	}

	return data
}
