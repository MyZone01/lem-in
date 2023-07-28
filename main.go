package main

import (
	"fmt"
	"lemin/lib"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		numberOfAnts, antFarm, fileContent, shouldReturn := lib.ParseFile(args[1])
		fmt.Println(fileContent)
		if shouldReturn {
			return
		}
		paths := lib.FindPaths(antFarm)

		lib.MoveAnts(paths, numberOfAnts)
	} else {
		fmt.Println("[USAGE] : go run . filename")
	}
}
