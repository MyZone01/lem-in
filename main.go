package main

import (
	"fmt"
	"lemin/utils"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		numberOfAnts, antFarm, shouldReturn := utils.ParseFile(args[1])
		if shouldReturn {
			return
		}

		fmt.Println("Number of Ants :", numberOfAnts)
		fmt.Println("Ant Farm :", antFarm)
	}
}
