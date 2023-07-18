package utils

import (
	"fmt"
	"lemin/model"
)

func MoveAnts(paths []model.Path, numberOfAnts int) {
	j := 0
	for i := 0; i < numberOfAnts; i++ {
		ant := model.Ant{
			Name: fmt.Sprintf("%d", i),
		}
		a := len(paths[j].Ants)
		b := len(paths[j].Rooms)

		d := len(paths[(j+1)%len(paths)].Ants)
		e := len(paths[(j+1)%len(paths)].Rooms)

		if a+b > e+d {
			paths[j+1].Ants = append(paths[j+1].Ants, ant)
			j = (j + 1) % len(paths)
		} else {
			paths[j].Ants = append(paths[j].Ants, ant)
		}

	}

	moves := [][]string{}
	for i := 0; i < numberOfAnts*len(paths); i++ {
		moves = append(moves, make([]string, numberOfAnts))
	}

	for _, path := range paths {
		for k := 0; k < len(path.Ants); k++ {
			_moves := []string{}
			for l := 0; l < len(path.Rooms); l++ {
				_moves = append(_moves, fmt.Sprintf("L%v-%v ", path.Ants[k].Name, path.Rooms[l].Name))
			}
			moves = append(moves, _moves)
		}
	}

	for _, v := range moves {
		fmt.Println(v)
	}
}
