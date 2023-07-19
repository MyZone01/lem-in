package utils

import "lemin/model"

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
	return StringToRoom(paths, antFarm)
}


