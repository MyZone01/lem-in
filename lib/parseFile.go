package lib

import (
	"bufio"
	"errors"
	"fmt"
	"lemin/models"
	"os"
	"strconv"
	"strings"
)

func ParseFile(fileName string) (int, models.AntFarm, string, bool) {

	file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("File Not Found")
			return 0, models.AntFarm{}, "", true
		}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	fileContent := ""

	for fileScanner.Scan() {
		str := strings.TrimRight(fileScanner.Text(), " ")
		str = strings.TrimLeft(str, " ")
		lines = append(lines, str)
		fileContent += fileScanner.Text() + "\n"
	}

	if len(lines) <= 0 {
	 	return 0, models.AntFarm{}, "", true
	}

	_numberOfAnts := lines[0]
	numberOfAnts, err := strconv.Atoi(_numberOfAnts)
	if err != nil || numberOfAnts <= 0 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		return 0, models.AntFarm{}, "", true
	}

	antFarm, err := GetAntFarmInfos(lines[1:])
	if err != nil {
		fmt.Println(err.Error())
		return 0, models.AntFarm{}, "", true
	}

	if antFarm.Start.Name == "" || antFarm.End.Name == "" {
		fmt.Println("ERROR: invalid data format, Missing start or end room")
		return 0, models.AntFarm{}, "", true
	}

	return numberOfAnts, antFarm, fileContent, false
}

func GetAntFarmInfos(lines []string) (models.AntFarm, error) {
	antFarm := models.AntFarm{}
	allRooms := map[string]models.Room{}
	if len(lines) > 0 {
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			if line == "" {
				continue
			}
			if strings.HasPrefix(line, "##") {
				if line == "##start" {
					if i+1 != len(lines) {
						i++
						line = lines[i]
						room, err := GetRoom(line)
						if err {
							return antFarm, errors.New("ERROR: invalid data format, Room bad formatted")
						}
						antFarm.Start = room
						allRooms[room.Name] = room
					} else {
						return antFarm, errors.New("ERROR: invalid data format, no start room found")
					}
				} else if line == "##end" {
					if i+1 != len(lines) {
						i++
						line = lines[i]
						room, err := GetRoom(line)
						if err {
							return antFarm, errors.New("ERROR: invalid data format, Room bad formatted")
						}
						antFarm.End = room
						allRooms[room.Name] = room
					} else {
						return antFarm, errors.New("ERROR: invalid data format, no start room found")
					}
				}
			} else {
				if !strings.HasPrefix(line, "#") {
					if strings.Contains(line, "-") {
						link, err := GetLink(line)
						if err {
							return antFarm, errors.New("ERROR: invalid data format, link bad formatted")
						}
						_, fromExist := allRooms[link.From]
						_, toExist := allRooms[link.To]
						if fromExist && toExist {
							if !haveDuplicate(link, antFarm.Links) {
								antFarm.Links = append(antFarm.Links, link)
							}
						} else {
							return antFarm, errors.New("ERROR: invalid data format, Link bad formatted")
						}
					} else {
						room, err := GetRoom(line)
						if err {
							return antFarm, errors.New("ERROR: invalid data format, Room bad formatted")
						}
						allRooms[room.Name] = room
					}
				}
			}
		}
		antFarm.Rooms = allRooms
		return antFarm, nil
	}
	return antFarm, errors.New("ERROR: invalid data format, Room bad formatted")
}
