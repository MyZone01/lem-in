package main

import (
	"fmt"
	"lemin/utils"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		numberOfAnts, antFarm, shouldReturn := utils.ParseFile(args[1])
		if shouldReturn {
			return
		}
		paths := utils.FindPaths(antFarm)
		
		utils.MoveAnts(paths, numberOfAnts)
	} else {
		fmt.Println("[USAGE] : go run . filename")
	}
}
