package utils

import (
	"fmt"
	"lemin/model"
)

func MoveAnt(paths []model.Path, numberOfAnts int) {
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
	for _, path := range paths {
		for k := 0; k < len(path.Ants); k++ {
			for l := 0; l < len(path.Rooms); l++ {
				move := fmt.Sprintf("L%v-%v\n", path.Ants[k].Name, path.Rooms[l].Name)
				if len(moves) < k {
					_move := []string{move}
					moves = append(moves, _move)
				} else {
					_move := []string{move}
					moves = append(moves, _move)
				}
			}
		}
	}

	// // activeAnts := []Ant{}
	// // for i := 0; i < numberOfAnts; i++ {
	// // fmt.Println("Ants: ", v.Ants)
	// // fmt.Println("Rooms: ", v.Rooms)
	// // k := 0
	// // l := 0
	// activeAnts := []model.Ant{}
	// for i := 0; i < numberOfAnts; i++ {
	// 	for _, path := range paths {
	// 		activeAnts = append(activeAnts, path.Ants[ant.CurrentRoom])
	// 	}
	// 	for _, ant := range activeAnts {
	// 		fmt.Printf("L%v-%s", ant.Name, ant.CurrentRoom)
	// 	}
	// }
	// // }
}
