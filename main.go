package main

import (
	"fmt"
)

type Coordinate struct {
	X int
	Y int
}

type TreasureHunt struct {
	CurrentPosition Coordinate
	TreasureHunt    Coordinate
	Map             [][]string
	IsGameOver      bool
	IsWinning       bool
}

func (th *TreasureHunt) newGame() {
	th.Map = [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", "X", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}
	th.TreasureHunt = Coordinate{
		X: 4,
		Y: 5,
	}
	th.IsGameOver = false
	th.IsWinning = false
	th.CurrentPosition = Coordinate{
		X: 4,
		Y: 1,
	}
}

func (th TreasureHunt) printMap() {
	for _, row := range th.Map {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println(" ")
	}
}

func (th TreasureHunt) addPossibleTreasureHunt() {
	x := th.CurrentPosition.X - 1
	var possibleRow []int

	for x > 0 || th.Map[x][th.CurrentPosition.Y] == "." {
		if th.Map[x][th.CurrentPosition.Y+1] == "." {
			possibleRow = append(possibleRow, x)
		}
		x--
	}

	colSize := len(th.Map[0])
	rowSize := len(th.Map)

	for _, row := range possibleRow {
		col := th.CurrentPosition.Y + 1

		for col < colSize && th.Map[row][col] != "#" {
			x := row + 1
			for x < rowSize && th.Map[x][col] == "." {
				th.Map[x][col] = "$"
				x++
			}
			col++
		}
	}
}

func (th *TreasureHunt) checkGameIsOver(direction string) {
	if direction == "north" && th.Map[th.CurrentPosition.X][th.CurrentPosition.Y+1] == "#" {
		th.IsGameOver = true
	} else if direction == "east" && th.Map[th.CurrentPosition.X+1][th.CurrentPosition.Y] == "#" {
		th.IsGameOver = true
	} else if direction == "south" {
		th.IsGameOver = true
		th.IsWinning = th.CurrentPosition.X == th.TreasureHunt.X && th.CurrentPosition.Y == th.TreasureHunt.Y
	}
}

func (th *TreasureHunt) move(direction string, steps int) string {
	if direction == "north" {
		if (th.CurrentPosition.X-steps) < 0 || th.Map[th.CurrentPosition.X-steps][th.CurrentPosition.Y] == "#" {
			return "Game over, couldn't find the treasure hunt"
		}

		th.Map[th.CurrentPosition.X][th.CurrentPosition.Y] = "."
		th.CurrentPosition.X = th.CurrentPosition.X - steps
		th.Map[th.CurrentPosition.X][th.CurrentPosition.Y] = "X"
	} else if direction == "east" {
		if (th.CurrentPosition.Y+steps) > len(th.Map[0]) || th.Map[th.CurrentPosition.X][th.CurrentPosition.Y+steps] == "#" {
			return "Game over, couldn't find the treasure hunt"
		}

		th.Map[th.CurrentPosition.X][th.CurrentPosition.Y] = "."
		th.CurrentPosition.Y = th.CurrentPosition.Y + steps
		th.Map[th.CurrentPosition.X][th.CurrentPosition.Y] = "X"
	} else {
		if (th.CurrentPosition.X+steps) > len(th.Map) || th.Map[th.CurrentPosition.X+steps][th.CurrentPosition.Y] == "#" {
			return "Game over, couldn't find the treasure hunt"
		}

		th.Map[th.CurrentPosition.X][th.CurrentPosition.Y] = "."
		th.CurrentPosition.X = th.CurrentPosition.X + steps
		th.Map[th.CurrentPosition.X][th.CurrentPosition.Y] = "X"
	}

	return ""
}

func main() {
	var th TreasureHunt
	th.newGame()

	printHeader("current map")
	th.printMap()

	printHeader("map with possible treasure hunt ($)")
	th.addPossibleTreasureHunt()
	th.printMap()

	directions := []string{"north", "east", "south"}
	var error string

	for _, direction := range directions {
		var steps int
		fmt.Println("please type in the number of steps to", direction)
		fmt.Scanln(&steps)
		if steps == 0 {
			fmt.Println("Game over, you enter an incorrect input")
		}
		error = th.move(direction, steps)
		if error != "" {
			fmt.Println(error)
			break
		}

		printHeader("current map")
		th.printMap()
		th.checkGameIsOver(direction)

		if th.IsGameOver {
			if th.IsWinning {
				fmt.Println("Congratulation, you found the treasure hunt")
			} else {
				fmt.Println("Game over, you couldn't move anymore")
			}
			break
		}
	}
}

func printHeader(text string) {
	fmt.Println("")
	fmt.Println(text)
	fmt.Println("===========")
}
