package utils

import (
	"lemin/model"
)

func findPathsRecursive(currentRoom string, path []string, visited map[string]bool, tunnels []model.Link, endRoom string, paths *[][]string) {
	visited[currentRoom] = true

	path = append(path, currentRoom)

	if currentRoom == endRoom {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*paths = append(*paths, newPath)
	} else {
		for _, tunnel := range tunnels {
			if tunnel.From == currentRoom && !visited[tunnel.To] {
				findPathsRecursive(tunnel.To, append([]string{}, path...), visited, tunnels, endRoom, paths)
			} else if tunnel.To == currentRoom && !visited[tunnel.From] {
				findPathsRecursive(tunnel.From, append([]string{}, path...), visited, tunnels, endRoom, paths)
			}
		}
	}

	visited[currentRoom] = false
}
func FindPaths(antFarm model.AntFarm) []model.Path {
	visited := make(map[string]bool)
	paths := [][]string{}

	// Fonction récursive pour trouver les chemins
	var findPathsRecursive func(currentRoom string, path []string)
	findPathsRecursive = func(currentRoom string, path []string) {
		// Marquer la salle actuelle comme visitée
		visited[currentRoom] = true

		// Ajouter la salle actuelle au chemin
		path = append(path, currentRoom)

		// Si la salle actuelle est la salle d'arrivée, ajouter le chemin complet à la liste des chemins
		if currentRoom == antFarm.End.Name {
			newPath := make([]string, len(path))
			copy(newPath, path)
			paths = append(paths, newPath)
		} else {
			// Parcourir les tunnels sortants de la salle actuelle
			for _, tunnel := range antFarm.Links {
				if tunnel.From == currentRoom && !visited[tunnel.To] {
					// Appel récursif pour la salle de destination du tunnel
					findPathsRecursive(tunnel.To, path)
				} else if tunnel.To == currentRoom && !visited[tunnel.From] {
					// Appel récursif pour la salle de destination du tunnel
					findPathsRecursive(tunnel.From, path)
				}
			}
		}

		// Marquer la salle actuelle comme non visitée après avoir exploré toutes les options
		visited[currentRoom] = false
	}

	// Appeler la fonction récursive avec la salle de départ et un chemin vide
	findPathsRecursive(antFarm.Start.Name, []string{})
	RangePaths(paths)
	return StringToRoom(paths, antFarm.Rooms)
}

func RangePaths(Paths [][]string) {
	for i := 0; i < len(Paths)-1; {
		if len(Paths[i]) > len(Paths[i+1]) {
			tmp := Paths[i]
			Paths[i] = Paths[i+1]
			Paths[i+1] = tmp
			i = 0
		} else {
			i++
		}

	}
}

func StringToRoom(Paths [][]string, Rooms map[string]model.Room) []model.Path {
	var PathsRoom []model.Path

	for _, path := range Paths {
		_rooms := []model.Room{}
		for _, room := range path {
			_rooms = append(_rooms, Rooms[room])
		}
		PathsRoom = append(PathsRoom, model.Path{Rooms: _rooms})
	}
	return PathsRoom
}
