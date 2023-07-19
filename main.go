package main

import (
	"bufio"
	"fmt"

	//"fmt"

	"lemin/utils"
	"log"
	"os"
)

func main() {
	A := os.Args
////////////////////////
	if len(A) > 1 {
		readFile, err := os.Open(A[1])
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var lines []string

		for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
		}
//////////////////////
		if len(lines) > 0 {

		NumberOfAnts := lines[0]

		if NumberOfAnts == "0" {
			fmt.Println("No Ant")
			return
		}
			//Recup The StartRoom and the EndRoom

		Start, err := utils.ReturnStart(lines)

		if err != nil {
			fmt.Println(err)
			return
		}

		End, err1 := utils.ReturnEnd(lines)

		if err1 != nil {
			fmt.Println(err1)
			return
		}

		//Recup the Links

		Links, err2 := utils.GetLink(lines)

		if err2 != nil {
			fmt.Println(err2)
			return
		}

		

		//Recup The Rooms
	
		Rooms, err := utils.GetRooms(lines)

		if err != nil {
			fmt.Println(err)
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		str := utils.FindPaths(Start.Name,End.Name,Links,Rooms)
		//RangePaths(str)
		// fmt.Println("First Paths : ", First)
		for _, val := range str {
			fmt.Println(val)

		}
		}
	}
}
