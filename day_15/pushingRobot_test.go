package main

import (
	"strings"
	"testing"
)

func TestMoveRobot(t *testing.T) {
	warehouse := ParseWarehouseMap(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	robotMove := '<'

	warehouse_got := warehouse.MoveRobot(robotMove)
	warehouse_expected := ParseWarehouseMap(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobot(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

	robotMove = '^'
	warehouse_got = warehouse.MoveRobot(robotMove)
	warehouse_expected = ParseWarehouseMap(`########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobot(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

	robotMove = '>'
	warehouse_got = warehouse.MoveRobot(robotMove)
	warehouse_expected = ParseWarehouseMap(`########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobot(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

	robotMove = '>' // this time two objects were pushed
	warehouse_got = warehouse.MoveRobot(robotMove)
	warehouse_expected = ParseWarehouseMap(`########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobot(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

	robotMove = '>'
	warehouse_got = warehouse.MoveRobot(robotMove)
	warehouse_expected = ParseWarehouseMap(`########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobot(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

	robotMove = 'v'
	warehouse_got = warehouse.MoveRobot(robotMove)
	warehouse_expected = ParseWarehouseMap(`########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobot(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

}

func TestMoveRobotSequence(t *testing.T) {
	warehouse := ParseWarehouseMap(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	robotMoves := []rune("<^^>>>vv<v>>v<<")

	warehouse_got := warehouse.MoveRobotSequence(robotMoves)
	warehouse_expected := ParseWarehouseMap(`########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobotSequence(%q) == %q, want %q", robotMoves, warehouse_got, warehouse_expected)
	}

}

func TestMoveRobotSequenceBigger(t *testing.T) {
	warehouse := ParseWarehouseMap(`##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########`)
	sequenceStr := `<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`
	robotMoves := []rune(strings.ReplaceAll(sequenceStr, "\n", ""))
	warehouse_got := warehouse.MoveRobotSequence(robotMoves)
	warehouse_expected := ParseWarehouseMap(`##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobotSequence(%q) == %q, want %q", robotMoves, warehouse_got, warehouse_expected)
	}

}

func TestSumBoxCoordinates(t *testing.T) {
	warehouse := ParseWarehouseMap(`########
#...O..
#......`)
	got := warehouse.SumBoxCoordinates()
	expected := 104
	if got != expected {
		t.Errorf("SumBoxCoordinates() == %d, want %d", got, expected)
	}

	warehouse = ParseWarehouseMap(`########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`)
	got = warehouse.SumBoxCoordinates()
	expected = 2028
	if got != expected {
		t.Errorf("SumBoxCoordinates() == %d, want %d", got, expected)
	}

}

func TestScaleUp(t *testing.T) {
	warehouse := ParseWarehouseMap(`#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######`)
	got := warehouse.ScaleUp()
	expected := ParseWarehouseMap(`##############
##......##..##
##..........##
##....[][]@.##
##....[]....##
##..........##
##############`)
	if got.Equals(expected) == false {
		t.Errorf("ScaleUp() == %q, want %q", got, expected)
	}
}

func TestMoveRobotPart2(t *testing.T) {
	warehouse := ParseWarehouseMap(`#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######`)
	warehouse.ScaleUp()
	robotMove := '<'
	warehouse_got := warehouse.MoveRobotPart2(robotMove)
	expected := ParseWarehouseMap(`##############
##......##..##
##..........##
##...[][]@..##
##....[]....##
##..........##
##############`)
	if warehouse_got.Equals(expected) == false {
		t.Errorf("MoveRobotPart2(%q) == %q, want %q", robotMove, warehouse_got, expected)
	}

	robotMove = 'v'
	warehouse_got = warehouse.MoveRobotPart2(robotMove)
	expected = ParseWarehouseMap(`##############
##......##..##
##..........##
##...[][]...##
##....[].@..##
##..........##
##############`)
	if warehouse_got.Equals(expected) == false {
		t.Errorf("MoveRobotPart2(%q) == %q, want %q", robotMove, warehouse_got, expected)
	}
}

func TestMoveRobotPart2_2(t *testing.T) {
	warehouse := ParseWarehouseMap(`##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##.....@....##
##############`)
	robotMove := '^'
	boxLocation := Location{warehouse.robotLocation.x, warehouse.robotLocation.y - 1}
	got := warehouse.IsPossibleToMoveObjectVertically(boxLocation, moves[robotMove])
	expected := true
	if got != expected {
		t.Errorf("IsPossibleToMoveObjectVertically(%q) == %t, want %t", robotMove, got, expected)
	}

	warehouse_got := warehouse.MoveRobotPart2(robotMove)

	warehouse_expected := ParseWarehouseMap(`##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobotPart2(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

	robotMove = '^'
	boxLocation = Location{warehouse.robotLocation.x, warehouse.robotLocation.y - 1}
	got = warehouse.IsPossibleToMoveObjectVertically(boxLocation, moves[robotMove])
	expected = false
	if got != expected {
		t.Errorf("IsPossibleToMoveObjectVertically(%q) == %t, want %t", robotMove, got, expected)
	}

	warehouse = ParseWarehouseMap(`##############
##......##..##
##...[][]...##
##...@[]....##
##..........##
##..........##
##############`)
	robotMove = '^'
	warehouse_got = warehouse.MoveRobotPart2(robotMove)
	warehouse_expected = ParseWarehouseMap(`##############
##...[].##..##
##...@.[]...##
##....[]....##
##..........##
##..........##
##############`)
	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobotPart2(%q) == %q, want %q", robotMove, warehouse_got, warehouse_expected)
	}

	warehouse = ParseWarehouseMap(`####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##...[].......[]..##
##[]##....[]......##
##[]......[]..[]..##
##..[][]..@[].[][]##
##........[]......##
####################`)

	warehouse_got = warehouse.MoveRobotPart2('^')

	warehouse_expected = ParseWarehouseMap(`####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##...[]...[]..[]..##
##[]##....[]......##
##[]......@...[]..##
##..[][]...[].[][]##
##........[]......##
####################`)

	if warehouse_got.Equals(warehouse_expected) == false {
		t.Errorf("MoveRobotPart2(%q) == %q, want %q", '^', warehouse_got, warehouse_expected)
	}

}

func TestIsPossibleToMoveObjectHorizontally(t *testing.T) {
	// state before step 310
	warehouse := ParseWarehouseMap(`####################
##[]..[]......[][]##
##[]........@..[].##
##..........[][][]##
##...........[][].##
##..##[]..[]......##
##...[]...[]..[]..##
##.....[]..[].[][]##
##........[]......##
####################`)

	robotMove := 'v'
	boxLocation := Location{warehouse.robotLocation.x, warehouse.robotLocation.y + 1}
	got := warehouse.IsPossibleToMoveObjectVertically(boxLocation, moves[robotMove])
	expected := true
	if got != expected {
		t.Errorf("IsPossibleToMoveObjectVertically(%q) == %t, want %t", robotMove, got, expected)
	}

	warehouse = ParseWarehouseMap(`####################################################################################################
##............[]..[]....[]............[]............[]....[][]......[]....[]..[]..[]....[]....[]..##
##........[]..[]....##..[]........[]##..##..##............[]..[]##[]..[]##[]##....[]..[]..........##
##..................[]....##[]............[]......[][]....[]....[]....[]##........[]..............##
##..[]..[]............[]##..##....##..##..................[]......[][][]..[]....[]....[]........[]##
##....[]............[]..............[]....[]..##[]............[][]..[]..[]##....[]..[][]....[]....##
##[]..[]..........[]..[]............[]........[]....[][]......##..##......[]####[][]##[]..........##
##[]....[]........[][][][]......##..[]..[]..[]##..........####..........[]##..####....[]..........##
##[]....[]..[]........[]..##....[]..[]..[][]##[]..##....##..[]........[]....[]..##..........[]....##
####....[][][][]....[]....##[]..##..[]........[]..............................[][]..........[]....##
##[]..[]##..................[][]....[]....##....[]..........[]....[]......[]##..........[]........##
##....##[]##[]##......##........[]......[][][][]....[]..........##[]..............[]..............##
##....................[]......[]..........##[][][][]##[]....##....[]....##[]..................[]..##
####....[]..[][]..[]............##......##[][][]..[][]....[]..[]..[]..[][]##..##[][]....[]..[][]..##
##........[]..[]....[][]..[]..[][][]....[][]......[]......[]....[]..[]..[]........[]......[]......##
##[]....##[][]..................[]##............[][][]........[]..........[]..[][]..##..[]..####..##
##..[]..[][][]..........[]........[]......##[]..[]..[]..##..[]..[]....................[]......[][]##
##..[][]..[][]......[]......[]....##................[]..[]......[]..[]......##[][]..[]........[]..##
##[][]......[]..[][]..##[]..........[][][]..##[]##..[]..##............[][]..........[][]##....[]..##
##....[]............[]..........##......##[][]..[]....[]..........[][][]..........[]..##[][]......##
##..........##[]....[]....##......[][][]..##.[].[]....[]..[]##......##[][]..[]............[]....####
##....[][]..[][][]##..[]..........[][]....##....[]............[][]..[]........[]......[]..........##
##..[]..[][]............................##[][]............[]..[]......[][]..........##..[]##......##
##............[][]..........[][][]..[]....[]##..........[][]....[]##....[]....[]..[][]............##
##..##[]........[]........[]....##[]....[]............##....##......##............[]..............##
##..##....[]..[]..[][]......[]..[]##..[]......[]..##....[]......[][][][][]......##..[][][]##......##
####[]....[]..[]........##[]........[][]##[]..##[]............[]..........[]....[]......##[]......##
##......[]......................[]..##.[][].....##......##..[]..[]..[]......................##..####
##[]##[]..............[]....[]..........@.....##[]..[]..##[][][][]##..[]..[]..[]..[]..............##
####[]..[]....[][]......##..[]....##..[].....[].[][][]..[]..[]##..........[]..[]....##..[]..[]..[]##
##....[][]......[]....##[]..[]..##..[][][][][]......[][]................##....[][]..##......[]..####
##[]..##..##....##..[]..[]......[]....[]..##[]....[]..[]..[][]............[]..[][]..[]......[]..####
##..[]..[][]......[][]..................##........[]..[][]..[][]....[]##[]..##[]................####
##..............[]..............[][]##......##........##..[][][]##..##................##..[]......##
##......[]..[]..##..[]..[]..................[][][]..[]..[]......[]........[]........##......[]##..##
##........[]..[]..............##..............##....[]####..[]..[]..........[]......[]..[]......[]##
##..[]..[]..[]....[]##..[]..[]......####[]##[]......[]..........##[][]##[]..[]........[]....[]....##
##[]..[]..[]..##....[]......[]..[]..............##..[][]....##....................[]......[]..[]..##
##......[]......[]..........[]........####............##....[][]........####..[]....[][]..........##
##....[][]....[]....[]....[][][]..........##[]....[][]..[]..[][][][]..##..[]......[]..##..........##
##......[]......[]..[]..[]......[]..##[]..##..[]......[]##..##............[]....[][]....[]..####[]##
##......[]..[]##............[][]..##..[]..........[][]....[]..............##......##..[]......##..##
####[][]....##[]..........[][]............................##......[]##....[]##............[][][]..##
####[][][][]....[]..##....[]..[]........[]........[]....[]##..[]##........[]........[]..[]........##
##..........##............[]......[][][]..[]................[]........##[]..[][]..........[][]..[]##
##[]..##........[][]##........[][]..[]........##[]##................##..##..[][]..##..[]......[][]##
##..[]..........[]..[][]..........[][]......##[]....[][][]..[]......[][]........[]..[]####[]......##
####[]....[][]..[][]....##....[]..[]..[]..####[]............[]....##........[][][]........[][]....##
##[]....[]........[]..................[][]##..##............[][]......####..[]..............[]....##
####################################################################################################`)
	robotMove = '^'
	boxLocation = Location{warehouse.robotLocation.x, warehouse.robotLocation.y + 1}
	got = warehouse.IsPossibleToMoveObjectVertically(boxLocation, moves[robotMove])
	expected = false
	if got != expected {
		t.Errorf("IsPossibleToMoveObjectVertically(%q) == %t, want %t", robotMove, got, expected)
	}
}

// func TestMoveRobotSequcePart2(t *testing.T) {
// 	warehouse := ParseWarehouseMap(`####################
// ##....[]....[]..[]##
// ##............[]..##
// ##..[][]....[]..[]##
// ##....[]@.....[]..##
// ##[]##....[]......##
// ##[]....[]....[]..##
// ##..[][]..[]..[][]##
// ##........[]......##
// ####################`)

// 	sequenceStr := `<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
// vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
// ><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
// <<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
// ^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
// ^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
// >^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
// <><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
// ^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
// v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`
// 	robotMoves := []rune(strings.ReplaceAll(sequenceStr, "\n", ""))

// 	warehouse_got := warehouse.MoveRobotSequencePart2(robotMoves)
// 	warehouse_expected := ParseWarehouseMap(`####################
// ##[].......[].[][]##
// ##[]...........[].##
// ##[]........[][][]##
// ##[]......[]....[]##
// ##..##......[]....##
// ##..[]............##
// ##..@......[].[][]##
// ##......[][]..[]..##
// ####################`)

// 	if warehouse_got.Equals(warehouse_expected) == false {
// 		t.Errorf("MoveRobotSequencePart2(%q) == %q, want %q", robotMoves, warehouse_got, warehouse_expected)
// 	}
// }
