// // https://adventofcode.com/2024/day/9
package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseRawDisk(rawDisk string) []int {
	// parse raw disk data format to slices of file_id and free space (represented by '-1')
	diskBlocks := make([]int, 0)
	fileID := 0
	for i, char := range rawDisk {
		if i%2 == 0 { // this is file size (in terms of data blocks)
			fileSize := int(char - '0')
			for j := 0; j < fileSize; j++ {
				diskBlocks = append(diskBlocks, fileID)
			}
			fileID += 1
		} else { // this is free space
			freeSpaceSize := int(char - '0')
			for j := 0; j < freeSpaceSize; j++ {
				diskBlocks = append(diskBlocks, -1)
			}
		}
		if i == len(rawDisk)-1 {
			break
		}
	}

	return diskBlocks
}

func FindNextFreeSpace(diskBlocks []int, start int) int {
	// find the next free space in the disk blocks
	for i := start; i < len(diskBlocks); i++ {
		if diskBlocks[i] == -1 {
			return i
		}
	}
	return -1
}

func DeFragmentDisk(diskBlocks []int) []int {
	// compress disk blocks by moving the file data blocks from the end to the free spaces from the beginning
	left := 0
	right := len(diskBlocks) - 1
	for left < right {
		if diskBlocks[left] != -1 {
			left++ // move until it meet a free space
			continue
		}
		if diskBlocks[right] == -1 {
			right-- // move until it meet a file data block
			continue
		}
		// now left points to a free space and right points to a file data block, swap their content
		// fmt.Println("swap content:", left, right)
		diskBlocks[left], diskBlocks[right] = diskBlocks[right], diskBlocks[left]
	}

	// now move left all the way to the end (when it meet a free space)
	for diskBlocks[left] != -1 && left < len(diskBlocks) {
		left++
	}

	return diskBlocks[:left]
}

func main() {
	rawDisk := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		rawDisk += line
	}

	fmt.Println("original input: ", rawDisk)

	// parse raw disk data format to slices of file_id and free space (represented by '-1')
	diskBlocksBeforeDefragmentation := ParseRawDisk(rawDisk)

	fmt.Printf("disk blocks before defragmentation: %v \n", diskBlocksBeforeDefragmentation)

	defragmendDiskBlocks := DeFragmentDisk(diskBlocksBeforeDefragmentation)

	fmt.Printf("disk blocks after defragmentation: %v \n", defragmendDiskBlocks)

	// calculate checksum
	checksum := 0
	for i, fileID := range defragmendDiskBlocks {
		checksum += i * fileID
	}
	// print checksum
	fmt.Println("checksum: ", checksum)
}
