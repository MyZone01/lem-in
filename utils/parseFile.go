package utils

import (
	"bufio"
	"fmt"
	"lemin/model"
	"log"
	"os"
	"strconv"
)

func ParseFile(fileName string) (int, model.AntFarm, bool) {
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
		return 0, model.AntFarm{}, true
	}
	_numberOfAnts := lines[0]
	if _numberOfAnts == "0" {
		fmt.Println("No Ant")
		return 0, model.AntFarm{}, true
	}

	numberOfAnts, err := strconv.Atoi(_numberOfAnts)
	if err != nil {
		fmt.Println(err)
		return 0, model.AntFarm{}, true
	}

	start, err := ReturnStart(lines)
	if err != nil {
		fmt.Println(err)
		return 0, model.AntFarm{}, true
	}
	end, err1 := ReturnEnd(lines)
	if err1 != nil {
		fmt.Println(err1)
		return 0, model.AntFarm{}, true
	}

	links, err2 := GetLink(lines)
	if err2 != nil {
		fmt.Println(err2)
		return 0, model.AntFarm{}, true
	}

	rooms, err := GetRooms(lines, start, end)
	if err != nil {
		fmt.Println(err)
		return 0, model.AntFarm{}, true
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
