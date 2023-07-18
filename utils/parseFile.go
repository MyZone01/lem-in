package utils

import (
	"bufio"
	"fmt"
	"lemin/model"
	"log"
	"os"
)

func ParseFile(fileName string) (string, model.AntFarm, bool) {
	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	if len(lines) <= 0 {
		return "", model.AntFarm{}, true
	}
	numberOfAnts := lines[0]
	if numberOfAnts == "0" {
		fmt.Println("No Ant")
		return "", model.AntFarm{}, true
	}

	start, err := ReturnStart(lines)
	if err != nil {
		fmt.Println(err)
		return "", model.AntFarm{}, true
	}
	end, err1 := ReturnEnd(lines)
	if err1 != nil {
		fmt.Println(err1)
		return "", model.AntFarm{}, true
	}

	links, err2 := GetLink(lines)
	if err2 != nil {
		fmt.Println(err2)
		return "", model.AntFarm{}, true
	}

	rooms, err := GetRooms(lines)
	if err != nil {
		fmt.Println(err)
		return "", model.AntFarm{}, true
	}

	validRoom, err3 := GetValidRooms(rooms, links)

	if err3 != nil {
		fmt.Println(err3)
	}

	antFarm := model.AntFarm{
		Start: start,
		End:   end,
		Rooms: validRoom,
		Links: links,
	}
	return numberOfAnts, antFarm, false
}
