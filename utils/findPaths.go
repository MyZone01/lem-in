package utils

import "lemin/model"

func FindPaths(antFarm model.AntFarm) []model.Path {
    visited := make(map[string]bool)
    paths := [][]string{}

    // Recursive function to find paths
    var findPathsRecursive func(currentRoom string, path []string)
    findPathsRecursive = func(currentRoom string, path []string) {
        visited[currentRoom] = true
        path = append(path, currentRoom)
        if currentRoom == antFarm.End.Name {
            newPath := make([]string, len(path))
            copy(newPath, path)

            // Check if newPath shares any room with existing paths
            newPathHasSharedRoom := false
            for _, existingPath := range paths {
                for _, room := range newPath {
                    if contains(existingPath, room) {
                        newPathHasSharedRoom = true
                        break
                    }
                }
                if newPathHasSharedRoom {
                    break
                }
            }

            // Only append newPath if it does not share any room with existing paths
            if !newPathHasSharedRoom {
                paths = append(paths, newPath)
            }
        } else {
            for _, tunnel := range antFarm.Links {
                if tunnel.From == currentRoom && !visited[tunnel.To] {
                    findPathsRecursive(tunnel.To, path)
                } else if tunnel.To == currentRoom && !visited[tunnel.From] {
                    findPathsRecursive(tunnel.From, path)
                }
            }
        }
        visited[currentRoom] = false
    }

    findPathsRecursive(antFarm.Start.Name, []string{})
    RangePaths(paths)
    return StringToRoom(paths, antFarm)
}

// Contains checks if a slice contains a string
func contains(slice []string, str string) bool {
    for _, item := range slice {
        if item == str {
            return true
        }
    }
    return false
}


