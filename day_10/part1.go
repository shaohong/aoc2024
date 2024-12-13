// generated by chatgpt
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Directions for moving up, down, left, right
var directions = [][2]int{
	{0, 1}, {1, 0},
	{0, -1}, {-1, 0},
}

func bfs(mapGrid [][]int, start [2]int) int {
	rows, cols := len(mapGrid), len(mapGrid[0])
	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true
	score := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		x, y := curr[0], curr[1]
		if mapGrid[x][y] == 9 {
			score++
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && ny >= 0 && nx < rows && ny < cols {
				next := [2]int{nx, ny}
				if !visited[next] && mapGrid[nx][ny] == mapGrid[x][y]+1 {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return score
}

func main() {
	// Read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var topographicMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, ch := range line {
			row[i], _ = strconv.Atoi(string(ch))
		}
		topographicMap = append(topographicMap, row)
	}

	// Calculate total score for all trailheads
	totalScore := 0
	for i := 0; i < len(topographicMap); i++ {
		for j := 0; j < len(topographicMap[0]); j++ {
			if topographicMap[i][j] == 0 { // Found a trailhead
				totalScore += bfs(topographicMap, [2]int{i, j})
			}
		}
	}

	fmt.Println(totalScore)
}