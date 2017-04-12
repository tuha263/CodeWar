package main

import (
	"fmt"
	"strconv"
)

const (
	// HEIGHT of matrix
	HEIGHT = 20
	// WEIGHT of matrix
	WEIGHT = 30
)

// Position of player
type Position struct {
	x int
	y int
}

//MoveLeft 1 value
func (curPos Position) MoveLeft() Position {
	return Position{curPos.x, curPos.y - 1}
}

//MoveRight 1 value
func (curPos Position) MoveRight() Position {
	return Position{curPos.x, curPos.y + 1}
}

//MoveUp 1 value
func (curPos Position) MoveUp() Position {
	return Position{curPos.x + 1, curPos.y}
}

//MoveDown 1 value
func (curPos Position) MoveDown() Position {
	return Position{curPos.x - 1, curPos.y}
}

//CanMoveLeft = true if can move left from current position
func (curPos Position) CanMoveLeft() bool {
	return curPos.y-1 >= 0
}

//CanMoveRight = true if can move left from current position
func (curPos Position) CanMoveRight() bool {
	return curPos.y+1 < WEIGHT
}

//CanMoveUp = true if can move left from current position
func (curPos Position) CanMoveUp() bool {
	return curPos.x-1 >= 0

}

//CanMoveDown = true if can move left from current position
func (curPos Position) CanMoveDown() bool {
	return curPos.y+1 < HEIGHT
}

var (
	firstLine      string
	matrix         [20]string
	playerPos      [4]Position
	myNum          int
	myPos          Position
	playerQuantity int
)

func main() {

	// Read data
	var firstLine string
	fmt.Scanln(&firstLine)
	var startLine int

	if len(firstLine) < 2 {
		playerQuantity, _ = strconv.Atoi(firstLine)
		fmt.Scanln(&myNum)
		matrix[0] = firstLine
		startLine = 1
	}

	for i := startLine; i < HEIGHT; i++ {
		fmt.Scanln(&matrix[i])
	}

	for i := 0; i < playerQuantity; i++ {
		fmt.Scan(&playerPos[i].x)
		fmt.Scan(&playerPos[i].y)
	}

	//AI

}

const (
	//MoveLeft output
	MoveLeft = "LEFT"
	//MoveRight output
	MoveRight = "RIGHT"
	//MoveUp output
	MoveUp = "UP"
	//MoveDown output
	MoveDown = "DOWN"
)

// GetMove return move of player
func GetMove() string {
	return MoveLeft
}

// GetStable get stable value of player
func GetStable(playerNum int) int {
	return playerNum*2 - 1
}

// GetUnStable get unstable value of player
func GetUnStable(playerNum int) int {
	return playerNum * 2
}

var travelFlag [HEIGHT][WEIGHT]bool
var travelList [HEIGHT * WEIGHT]Position

// DistanceToValue get distance to safe of player
func DistanceToValue(pos Position, value int, distance int) int {
	//init
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WEIGHT; j++ {
			travelFlag[i][j] = false
		}
	}
	travelList[0] = pos
	first := 0
	last := 0

	result := 0
	for first <= last {
		first++

	}

	return result
}

//CheckTravelFlag get travel flag of flag matrix
func CheckTravelFlag(pos Position) {
	return travelFlag[pos.x][pos.y]
}

// Travel a position
func Travel(pos Position, first int, last int) (int, int) {
	first++
	travelFlag[pos.x][pos.y] = true
	if pos.CanMoveLeft() && !CheckTravelFlag(pos.MoveLeft()) {
		last++
		travelList[last] = pos.MoveLeft()
	}

	if pos.CanMoveRight() && !CheckTravelFlag(pos.MoveRight()) {
		last++
		travelList[last] = pos.MoveRight()
	}

	if pos.CanMoveUp() && !CheckTravelFlag(pos.MoveUp()) {
		last++
		travelList[last] = pos.MoveUp()
	}

	if pos.CanMoveDown() && !CheckTravelFlag(pos.MoveDown()) {
		last++
		travelList[last] = pos.MoveDown()
	}

	return first, last
}

// GetMatrixValue get value at position of matrix
func GetMatrixValue(pos Position) int {
	return int(matrix[pos.x][pos.y] - '0')
}

// GetDistance get distance of two pos
func GetDistance(pos1, pos2 Position) int {
	return GetIntABS(pos1.x-pos2.x) + GetIntABS(pos1.y-pos2.y)
}

// GetIntABS get abs
func GetIntABS(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
