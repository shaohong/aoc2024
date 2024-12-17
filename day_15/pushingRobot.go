// https://adventofcode.com/2024/day/15

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Location struct {
	x, y int
}

const (
	EmptySymbol    = '.'
	ObjectSymbol   = 'O'
	RobotSymbol    = '@'
	WallSymbol     = '#'
	ObjectSymbolWL = '['
	ObjectSymbolWR = ']'
)

type WareHouse struct {
	warehouseMap  [][]rune
	robotLocation Location
}

func ParseWarehouseMap(input string) WareHouse {
	lines := strings.Split(input, "\n")
	warehouseMap := make([][]rune, len(lines))

	for i, line := range lines {
		warehouseMap[i] = []rune(line)
	}

	robotLocation := Location{}
	for y, row := range warehouseMap {
		for x, cell := range row {
			if cell == RobotSymbol {
				robotLocation = Location{x, y}
				break
			}
		}
		if robotLocation != (Location{}) {
			break
		}
	}

	return WareHouse{warehouseMap, robotLocation}
}

func (w WareHouse) String() string {
	var sb strings.Builder

	for _, row := range w.warehouseMap {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}

	return sb.String()
}

func (w WareHouse) Equals(other WareHouse) bool {
	if len(w.warehouseMap) != len(other.warehouseMap) {
		fmt.Println("Warehouse maps have different heights", len(w.warehouseMap), len(other.warehouseMap))
		return false
	}
	for i := range w.warehouseMap {
		if len(w.warehouseMap[i]) != len(other.warehouseMap[i]) {
			fmt.Println("Warehouse maps have different widths")
			return false
		}

		for j := range w.warehouseMap[i] {
			if w.warehouseMap[i][j] != other.warehouseMap[i][j] {
				fmt.Printf("Warehouse maps differ at (%d, %d)\n", i, j)
				return false
			}
		}
	}
	return true
}

func ParseRobotMoves(input string) []rune {
	robotMoves := []rune(strings.ReplaceAll(input, "\n", ""))

	return robotMoves
}

// Directions for moving up, down, left, right
var moves = map[rune][2]int{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

func (w *WareHouse) IsEmpty(location Location) bool {
	return w.warehouseMap[location.y][location.x] == EmptySymbol
}

func (w *WareHouse) GetSymbol(location Location) rune {
	return w.warehouseMap[location.y][location.x]
}

// move object at the given location along the moveDir
func (w *WareHouse) MoveObject(location Location, moveDir [2]int) bool {
	currentSymbol := w.GetSymbol(location)
	newLocation := Location{location.x + moveDir[0], location.y + moveDir[1]}

	switch w.GetSymbol(newLocation) {
	case WallSymbol: // can't move object to the wall
		return false

	case EmptySymbol: // move object to the empty new location
		w.warehouseMap[location.y][location.x] = EmptySymbol
		w.warehouseMap[newLocation.y][newLocation.x] = currentSymbol
		return true

	default: // recursively move the objects
		if w.MoveObject(newLocation, moveDir) {
			w.warehouseMap[location.y][location.x] = EmptySymbol
			w.warehouseMap[newLocation.y][newLocation.x] = currentSymbol
			return true
		} else {
			return false
		}
	}
}

func (w *WareHouse) UpdateRobotLocation(newLocation Location) {
	w.warehouseMap[w.robotLocation.y][w.robotLocation.x] = EmptySymbol
	w.warehouseMap[newLocation.y][newLocation.x] = RobotSymbol
	w.robotLocation = newLocation
}

// move the robot, return new position of the robot and the updated warehouse map
func (w *WareHouse) MoveRobot(move rune) WareHouse {

	moveDir := moves[move]
	newRobotLocation := Location{w.robotLocation.x + moveDir[0], w.robotLocation.y + moveDir[1]}

	// check if the new location is valid
	if newRobotLocation.x < 0 || newRobotLocation.x >= len(w.warehouseMap[0]) ||
		newRobotLocation.y < 0 || newRobotLocation.y >= len(w.warehouseMap) {
		return *w
	}

	// check if the new location is a wall
	if w.warehouseMap[newRobotLocation.y][newRobotLocation.x] == WallSymbol {
		return *w
	}

	switch w.GetSymbol(newRobotLocation) {
	case EmptySymbol:
		// move the robot
		w.UpdateRobotLocation(newRobotLocation)
	case ObjectSymbol:
		if w.MoveObject(newRobotLocation, moveDir) {
			// object moved away, move the robot to the new location
			w.UpdateRobotLocation(newRobotLocation)
		}
	}
	return *w
}

func (w *WareHouse) MoveRobotSequence(moves []rune) WareHouse {
	for _, move := range moves {
		w.MoveRobot(move)
	}
	return *w
}

func (w *WareHouse) SumBoxCoordinates() int {
	// The GPS coordinate of a box is equal to 100 times its distance from the top edge of the map plus its distance from the left edge of the map.
	// we return the sum of the GPS coordinates of all boxes in the warehouse
	var sum int = 0
	for y, row := range w.warehouseMap {
		for x, cell := range row {
			if cell == ObjectSymbol {
				sum += 100*y + x
			}
		}
	}

	return sum
}

func (w *WareHouse) SumBoxCoordinatesPart2() int {
	var sum int = 0
	for y, row := range w.warehouseMap {
		for x, cell := range row {
			if cell == ObjectSymbolWL {
				sum += 100*y + x
			}
		}
	}

	return sum
}

func (w *WareHouse) ScaleUp() WareHouse {
	/*To get the wider warehouse's map, start with your original map and, for each tile, make the following changes:

	  If the tile is #, the new map contains ## instead.
	  If the tile is O, the new map contains [] instead.
	  If the tile is ., the new map contains .. instead.
	  If the tile is @, the new map contains @. instead. */

	// create a new warehouse map with double the width
	newMap := make([][]rune, len(w.warehouseMap))
	for i, row := range w.warehouseMap {
		newMap[i] = make([]rune, 2*len(row))
		for j, cell := range row {
			switch cell {
			case WallSymbol:
				newMap[i][2*j] = WallSymbol
				newMap[i][2*j+1] = WallSymbol
			case ObjectSymbol:
				newMap[i][2*j] = ObjectSymbolWL
				newMap[i][2*j+1] = ObjectSymbolWR
			case EmptySymbol:
				newMap[i][2*j] = EmptySymbol
				newMap[i][2*j+1] = EmptySymbol
			case RobotSymbol:
				newMap[i][2*j] = RobotSymbol
				newMap[i][2*j+1] = EmptySymbol
				w.robotLocation = Location{2 * j, i}
			}
		}
	}

	w.warehouseMap = newMap

	return *w
}

func getPairLocation(symbol rune, location Location) Location {
	var pairLocation Location
	if symbol == ObjectSymbolWL {
		pairLocation = Location{location.x + 1, location.y}
	} else {
		pairLocation = Location{location.x - 1, location.y}
	}
	return pairLocation
}

func getPairSymbol(symbol rune) rune {
	if symbol == ObjectSymbolWL {
		return ObjectSymbolWR
	}
	return ObjectSymbolWL
}

// check if it's possible to move the object at the given location along the moveDir, in Part 2
// this mostly concerns the movement in vertica	directions
func (w *WareHouse) IsPossibleToMoveObjectVertically(location Location, moveDir [2]int) bool {
	if moveDir[1] == 0 {
		panic("moveDir[1] should be non-zero")
	}

	currentSymbol := w.GetSymbol(location)
	pairLocation := getPairLocation(currentSymbol, location)

	if currentSymbol == ObjectSymbolWR {
		// switch the location and pair so that the location always has '[' part of the box
		location, pairLocation = pairLocation, location
		currentSymbol = ObjectSymbolWL
	}

	newLocationL := Location{location.x + moveDir[0], location.y + moveDir[1]}
	newLocationR := Location{pairLocation.x + moveDir[0], pairLocation.y + moveDir[1]}

	if w.GetSymbol(newLocationL) == EmptySymbol && w.GetSymbol(newLocationR) == EmptySymbol {
		return true
	} else if w.GetSymbol(newLocationL) == WallSymbol || w.GetSymbol(newLocationR) == WallSymbol {
		return false
	} else {
		// have to be checked recursively
		if w.GetSymbol(newLocationL) == EmptySymbol {
			return w.IsPossibleToMoveObjectVertically(newLocationR, moveDir)
		} else if w.GetSymbol(newLocationR) == EmptySymbol {
			return w.IsPossibleToMoveObjectVertically(newLocationL, moveDir)
		} else {
			// both are not empty, check if the object can be moved
			return w.IsPossibleToMoveObjectVertically(newLocationL, moveDir) && w.IsPossibleToMoveObjectVertically(newLocationR, moveDir)
		}

	}
}

// recursively move box parts at the location, vertically
func (w *WareHouse) MoveObjectVerticallyPart2(location Location, moveDir [2]int) {
	// recursively move object parts vertically
	currentSymbol := w.GetSymbol(location)
	if currentSymbol != ObjectSymbolWL && currentSymbol != ObjectSymbolWR {
		fmt.Println("currentSymbol is ", string(currentSymbol), " at ", location)
		panic("currentSymbol should be [ or ]")
	}

	pairLocation := getPairLocation(currentSymbol, location)

	if currentSymbol == ObjectSymbolWR {
		// switch the location and pair so that the location is always the '[' part of the box
		location, pairLocation = pairLocation, location
		currentSymbol = ObjectSymbolWL
	}

	pairSymbol := getPairSymbol(currentSymbol)

	newLocationL := Location{location.x + moveDir[0], location.y + moveDir[1]}
	newLocationR := Location{pairLocation.x + moveDir[0], pairLocation.y + moveDir[1]}

	if w.GetSymbol(newLocationL) == EmptySymbol && w.GetSymbol(newLocationR) == EmptySymbol {

		// current object pair can be moved now
		fmt.Printf("now move %c at %v to %v\n", currentSymbol, location, newLocationL)
		fmt.Printf("now move %c at %v to %v\n", pairSymbol, pairLocation, newLocationR)

		w.warehouseMap[newLocationL.y][newLocationL.x] = ObjectSymbolWL
		w.warehouseMap[newLocationR.y][newLocationR.x] = ObjectSymbolWR
		w.warehouseMap[location.y][location.x] = EmptySymbol
		w.warehouseMap[pairLocation.y][pairLocation.x] = EmptySymbol
	} else {
		// find the boxes that need to be moved, push them into a stack

		if w.GetSymbol(newLocationL) != EmptySymbol {
			w.MoveObjectVerticallyPart2(newLocationL, moveDir)
		}
		w.warehouseMap[newLocationL.y][newLocationL.x] = ObjectSymbolWL
		w.warehouseMap[location.y][location.x] = EmptySymbol

		if w.GetSymbol(newLocationR) != EmptySymbol {
			w.MoveObjectVerticallyPart2(newLocationR, moveDir)
		}
		w.warehouseMap[newLocationR.y][newLocationR.x] = ObjectSymbolWR
		w.warehouseMap[pairLocation.y][pairLocation.x] = EmptySymbol
	}

}

// move object in the current location along the moveDir to the new location
func (w *WareHouse) MoveObjectPart2(location Location, moveDir [2]int) bool {
	if moveDir[1] == 0 {
		return w.MoveObject(location, moveDir)
	}

	if !w.IsPossibleToMoveObjectVertically(location, moveDir) {
		return false
	} else {
		w.MoveObjectVerticallyPart2(location, moveDir)
		return true
	}
}

// move robot in the resized/scaledup warehouse
func (w *WareHouse) MoveRobotPart2(move rune) WareHouse {
	// fmt.Println("MoveRobotPart2: ", string(move))
	moveDir := moves[move]
	newRobotLocation := Location{w.robotLocation.x + moveDir[0], w.robotLocation.y + moveDir[1]}

	// check if the new location is valid
	if newRobotLocation.x < 0 || newRobotLocation.x >= len(w.warehouseMap[0]) ||
		newRobotLocation.y < 0 || newRobotLocation.y >= len(w.warehouseMap) {
		return *w
	}

	// check if the new location is a wall
	if w.warehouseMap[newRobotLocation.y][newRobotLocation.x] == WallSymbol {
		return *w
	}

	switch w.GetSymbol(newRobotLocation) {
	case EmptySymbol:
		// move the robot
		w.UpdateRobotLocation(newRobotLocation)
	case ObjectSymbolWL, ObjectSymbolWR:
		// try to move the object
		if w.MoveObjectPart2(newRobotLocation, moveDir) {
			// object moved away, move the robot to the new location
			w.UpdateRobotLocation(newRobotLocation)
		}
	}
	return *w
}

func (w *WareHouse) MoveRobotSequencePart2(moves []rune) WareHouse {
	for i, move := range moves {
		step := i + 1
		fmt.Printf("%d/%d - Move %c:\n", step, len(moves), move)

		w.MoveRobotPart2(move)

		// if step >= 170 && step <= 173 {
		// 	fmt.Println(w.String())
		// }
	}
	return *w
}

func main() {

	fmt.Println("Pushing Robots")
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	input_parts := strings.Split(string(data), "\n\n")
	warehouse := ParseWarehouseMap(input_parts[0])
	robotMoves := ParseRobotMoves(input_parts[1])

	fmt.Println("Warehouse Map:")
	fmt.Println(warehouse.String())

	fmt.Println("Robot Moves:\n", string(robotMoves))

	warehouse.MoveRobotSequence(robotMoves)
	fmt.Println("Warehouse Map after robot moves:")
	fmt.Println(warehouse.String())

	fmt.Println("Sum of box GPS coordinates:\n", warehouse.SumBoxCoordinates())

	//----------------- Part 2 -----------------
	warehouse = ParseWarehouseMap(input_parts[0])
	warehouse.ScaleUp()
	warehouse.MoveRobotSequencePart2(robotMoves)
	fmt.Println("Sum of box GPS coordinates, Part2:\n", warehouse.SumBoxCoordinatesPart2())

}
