package utils

import "lemin/model"

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