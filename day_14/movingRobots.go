// https://adventofcode.com/2024/day/14
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Location struct {
	x int
	y int
}

type Velocity struct {
	vx int
	vy int
}

// a robot has a location and a velocity
type Robot struct {
	location Location
	velocity Velocity
}

func (r *Robot) String() string {
	return fmt.Sprintf("Robot at (%d, %d) with velocity (%d, %d)", r.location.x, r.location.y, r.velocity.vx, r.velocity.vy)
}

// move the robot for a given number of seconds, return the new location
func (r Robot) move(seconds int, xLimit int, yLimit int) (newLocation Location) {
	newLocation.x = (r.location.x + r.velocity.vx*seconds) % xLimit
	if newLocation.x < 0 {
		newLocation.x += xLimit
	}

	newLocation.y = (r.location.y + r.velocity.vy*seconds) % yLimit
	if newLocation.y < 0 {
		newLocation.y += yLimit
	}

	return
}

func ParseInput(input string) []Robot {
	robots := []Robot{}

	for _, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var x, y, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		robots = append(robots, Robot{Location{x, y}, Velocity{vx, vy}})
	}

	return robots
}

func PrintRobots(robots []Robot, xLimit int, yLimit int) {
	for y := 0; y < yLimit; y++ {
		for x := 0; x < xLimit; x++ {
			robotCount := 0
			for _, robot := range robots {
				if robot.location.x == x && robot.location.y == y {
					robotCount++
				}
			}
			if robotCount > 0 {
				fmt.Print(robotCount)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

// maps the robot to its quadrant, -1 if it's on the dividing line
// 0|1
// -+-
// 2|3
func GetRobotQudrant(robot Robot, xLimit int, yLimit int) int {

	if robot.location.x == xLimit/2 || robot.location.y == yLimit/2 {
		return -1
	}

	if robot.location.x < xLimit/2 {
		if robot.location.y < yLimit/2 {
			return 0
		} else {
			return 2
		}
	} else {
		if robot.location.y < yLimit/2 {
			return 1
		} else {
			return 3
		}
	}
}

// put robots into quadrants
func Quadrantize(robots []Robot, xLimit int, yLimit int) map[int]int {

	quadrants := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}

	for _, robot := range robots {
		qudrant := GetRobotQudrant(robot, xLimit, yLimit)
		if qudrant == -1 {
			continue
		}
		quadrants[qudrant] = quadrants[qudrant] + 1
	}

	return quadrants
}

func DetectContinuousRegion(robots []Robot, xLimit int, yLimit int) bool {

	// put the robots into the the 2d matrix and detect continues distribution of robots
	matrix := make([][]int, yLimit)

	for y := 0; y < yLimit; y++ {
		matrix[y] = make([]int, xLimit)
	}

	for _, robot := range robots {
		matrix[robot.location.y][robot.location.x] = 1
	}

	// now check for continous distribution of robots
	gridSize := 5
	densityThreshold := 0.8

	for _, robot := range robots {
		// see how many robots are in the region within the gridsize distance to the current robot
		robotCount := 0
		for y := robot.location.y - gridSize/2; y <= robot.location.y+gridSize/2; y++ {
			for x := robot.location.x - gridSize/2; x <= robot.location.x+gridSize/2; x++ {
				if y < 0 || y >= yLimit || x < 0 || x >= xLimit {
					continue
				}
				if matrix[y][x] == 1 {
					robotCount++
				}
			}
		}
		if float64(robotCount) >= float64(gridSize*gridSize)*densityThreshold {
			return true
		}
	}

	return false
}

func Part2(robots []Robot, xLimit int, yLimit int) {

	// a chrimas tree shaped pattern is formed by the robots

	seconds := 0
	for {
		newRobots := make([]Robot, 0, len(robots))
		seconds++

		for _, robot := range robots {
			newLocation := robot.move(seconds, xLimit, yLimit)
			newRobots = append(newRobots, Robot{newLocation, robot.velocity})
		}

		continousRegionDetected := DetectContinuousRegion(newRobots, xLimit, yLimit)

		if continousRegionDetected {
			fmt.Println("Continous region detected at", seconds, "seconds")
			PrintRobots(newRobots, xLimit, yLimit)
			break
		}
		if seconds%1000 == 0 {
			fmt.Println(seconds, "seconds passed")
		}
	}

}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	robots := ParseInput(string(data))
	for _, robot := range robots {
		fmt.Println(robot.String())
	}

	// for the small input
	// xLimit := 11
	// yLimit := 7

	// for larger input
	xLimit := 101
	yLimit := 103

	fmt.Println("number of robots:", len(robots))
	fmt.Println("Initial state:")
	PrintRobots(robots, xLimit, yLimit)

	numberOfSeconds := 100
	newRobots := make([]Robot, 0, len(robots))
	for _, robot := range robots {
		newLocation := robot.move(numberOfSeconds, xLimit, yLimit)
		newRobots = append(newRobots, Robot{newLocation, robot.velocity})
	}

	fmt.Println("number of robots:", len(newRobots))
	fmt.Println("After", numberOfSeconds, "seconds")
	PrintRobots(newRobots, xLimit, yLimit)

	qudrantsCount := Quadrantize(newRobots, xLimit, yLimit)
	fmt.Println("Quadrants count:", qudrantsCount)

	safetyFactor := 1
	for _, count := range qudrantsCount {
		safetyFactor *= count
	}
	fmt.Println("Safety factor:", safetyFactor)

	Part2(robots, xLimit, yLimit)
}
