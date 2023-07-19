package utils

import (
	"fmt"
	"lemin/model"
)

type antPosition struct {
	pathIdx int
	roomIdx int
}

func MoveAnts(paths []model.Path, numberOfAnts int) {
	currentPath := 0
	antPositions := make(map[string]*antPosition)
	for i := 0; i < numberOfAnts; i++ {
		nextPath := (currentPath + 1) % len(paths)
		ant := model.Ant{
			Name: fmt.Sprintf("%d", i),
		}
		currentPathValue := len(paths[currentPath].Ants) + len(paths[currentPath].Rooms)
		nextPathValue := len(paths[nextPath].Ants) + len(paths[nextPath].Rooms)

		if currentPathValue > nextPathValue {
			paths[nextPath].Ants = append(paths[nextPath].Ants, ant)
			currentPath = nextPath
		} else {
			paths[currentPath].Ants = append(paths[currentPath].Ants, ant)
		}
		i := len(paths[currentPath].Ants) - 1
		if i != 0 {
			i *= -1
		}
		antPositions[ant.Name] = &antPosition{currentPath, i}
	}

	for {
		finished := true
		for antName, pos := range antPositions {
			if pos.roomIdx >= 0 && pos.roomIdx < len(paths[pos.pathIdx].Rooms) {
				room := paths[pos.pathIdx].Rooms[pos.roomIdx]
				fmt.Printf("L%s-%s ", antName, room.Name)
				pos.roomIdx++
				finished = false
			} else {
				antPositions[antName].roomIdx++
			}
		}
		if finished {
			break
		} else {
			fmt.Println()
		}
	}
}
