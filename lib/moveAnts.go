package lib

import (
	"fmt"
	"lemin/models"
	"strings"
)

type antPosition struct {
	pathIdx int
	roomIdx int
}

func MoveAnts(paths []models.Path, numberOfAnts int) {
	if len(paths) != 0 {
		currentPath := 0
		antPositions := make(map[string]*antPosition)
		for i := 0; i < numberOfAnts; i++ {
			nextPath := (currentPath + 1) % len(paths)
			ant := models.Ant{
				Name: fmt.Sprintf("%d", i+1),
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

		moves := ""
		for {
			finished := true
			for antName, pos := range antPositions {
				if pos.roomIdx >= 0 && pos.roomIdx < len(paths[pos.pathIdx].Rooms) {
					room := paths[pos.pathIdx].Rooms[pos.roomIdx]
					moves += fmt.Sprintf("L%s-%s ", antName, room.Name)
					pos.roomIdx++
					finished = false
				} else {
					antPositions[antName].roomIdx++
				}
			}
			if finished {
				break
			} else {
				moves += "\n"
			}
		}

		fmt.Println(moves)
		fmt.Println("🐜🐜 We have :", len(strings.Split(moves, "\n"))-1, "turns")
	} else {
		fmt.Println("There is no path")
	}
}
