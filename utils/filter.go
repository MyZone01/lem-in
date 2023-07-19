package utils

import "lemin/model"

func FilterLinks(tab []model.Link) []model.Link {
	var bat []model.Link 

	for _, val := range tab {
		
		if !IsPresent(val, bat) {
			bat = append(bat, val)
		}
	}

	return bat
}

func IsPresent(Link model.Link, tab []model.Link) bool {
	for _, val := range tab {
		if (val.From == Link.From && val.To == Link.To) || (val.To == Link.From && val.From == Link.To) {
			return true
		}
	}
	return false
}