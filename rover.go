package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func processFile(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("Unable to read file")
	}
	reader := bufio.NewReader(file)
	var lenX, lenY, startX, startY int64
	var currentOrient, movements string

	n, err := fmt.Fscanf(reader, "%d %d\n", &lenX, &lenY)
	if n != 2 && err != nil {
		fmt.Println("Unable to read size of grid successfully")
	}
	for {
		n1, err1 := fmt.Fscanf(reader, "%d %d %s\n", &startX, &startY, &currentOrient)
		n2, err2 := fmt.Fscanf(reader, "%s\n", &movements)
		if (n1 != 3 || err1 != nil) || (n2 != 1 || err2 != nil) {
			if err1 == io.EOF || err2 == io.EOF {
				fmt.Println("End of File Reached")
			} else if n1 != 3 || n2 != 1 {
				fmt.Println(n1, n2)
				fmt.Println("Wrong input provided")
			}
			break
		}
		finalX, finalY, dir, err := moveRover(startX, startY, currentOrient, movements, lenX, lenY)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Final Pos: ", finalX, finalY, dir)
		}
	}
}

func moveOneStep(curX, curY *int64, curOrient int, lenX, lenY int64) error {
	if curOrient%2 == 0 {
		if curOrient > 1 {
			*curY -= 1
		} else {
			*curY += 1
		}
		if *curY < 0 || *curY > lenY {
			return fmt.Errorf("rover moved outside the grid from Y axis with y-coordinate: %v", *curY)
		}
	} else {
		if curOrient > 2 {
			*curX += 1
		} else {
			*curX -= 1
		}
		if *curX < 0 || *curY > lenX {
			return fmt.Errorf("rover moved outside the grid from X axis with x-coordinate: %v", *curX)
		}
	}
	return nil
}

func moveRover(curX, curY int64, currentOrient string, movements string, lenX, lenY int64) (int64, int64, string, error) {

	directionStrToInt := map[string]int{"N": 0, "E": 1, "S": 2, "W": 3}
	roverOrient, _ := directionStrToInt[currentOrient]
	for _, move := range movements {
		switch move {
		case 'L':
			roverOrient = roverOrient + 4 - 1
			roverOrient %= 4
		case 'M':
			err := moveOneStep(&curX, &curY, roverOrient, lenX, lenY)
			if err != nil {
				return 0, 0, "", err
			}
		case 'R':
			roverOrient += 1
			roverOrient %= 4
		}
		//fmt.Println("New Position after move", string(move), " : ", curX, curY, roverOrient)
	}
	directionIntToStr := map[int]string{0: "N", 1: "E", 2: "S", 3: "W"}
	currentDirection, _ := directionIntToStr[roverOrient]
	return curX, curY, currentDirection, nil
}
