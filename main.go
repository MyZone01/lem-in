package main

import (
	"bufio"
	"fmt"
	"lemin/model"
	"lemin/utils"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		readFile, err := os.Open(args[1])
		if err != nil {
			log.Fatal(err)
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var lines []string
		for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
		}

		if len(lines) > 0 {
			numberOfAnts := lines[0]
			if numberOfAnts == "0" {
				fmt.Println("No Ant")
				return
			}

			// Retrieve The StartRoom and the EndRoom
			start, err := utils.ReturnStart(lines)
			if err != nil {
				fmt.Println(err)
				return
			}
			end, err1 := utils.ReturnEnd(lines)
			if err1 != nil {
				fmt.Println(err1)
				return
			}
			fmt.Println("Start :", start)
			fmt.Println("End :", end)

			// Retrieve the links
			links, err2 := utils.GetLink(lines)
			if err2 != nil {
				fmt.Println(err2)
				return
			}
			fmt.Println("Links :", links)

			// Retrieve The rooms
			rooms, err := utils.GetRooms(lines)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Rooms :", rooms)

			// Get Valid Rooms
			validRoom, err3 := utils.GetValidRooms(rooms, links)

			if err3 != nil {
				fmt.Println(err3)
			}

			fmt.Println("Valid Rooms : ", validRoom)

			antFarm := model.AntFarm{
				Start: start,
				End: end,
				Rooms: rooms,
				Links: links,
			}

			fmt.Println("Ant Farm :", antFarm)
		}

	}

}
