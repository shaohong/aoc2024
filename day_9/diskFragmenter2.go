// https://adventofcode.com/2024/day/9
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type DiskSegment struct {
	fileID int // -1 means free space
	size   int
}

func ParseRawDisk(rawDisk string) (diskBlocks []DiskSegment) {
	diskBlocks = make([]DiskSegment, 0)

	// parse the raw disk format into a linked list where each node is either file (id= fileID, size=fileSize) or free space (id = -1, size=freespaceSize)

	fileID := 0
	for i, char := range rawDisk {
		if i%2 == 0 { // this is file size (in terms of data blocks)
			fileSize := int(char - '0')
			diskBlocks = append(diskBlocks, DiskSegment{fileID, fileSize})
			fileID += 1
		} else { // this is free space
			freeSpaceSize := int(char - '0')
			diskBlocks = append(diskBlocks, DiskSegment{-1, freeSpaceSize})
		}
	}

	return diskBlocks
}

func DefragmentWithWholeFileMove(diskBlocks []DiskSegment) []DiskSegment {
	newDiskSegments := make([]DiskSegment, len(diskBlocks))
	_ = copy(newDiskSegments, diskBlocks)

	rightOffset := 1 // check each fileSegment from the right side
	for rightOffset < len(newDiskSegments) {
		// make sure right offet points to a file segment
		if newDiskSegments[len(newDiskSegments)-rightOffset].fileID == -1 {
			rightOffset += 1
			continue
		}
		currentFileSegment := newDiskSegments[len(newDiskSegments)-rightOffset]

		// find the first freeSpace segment from the left, that can accomodate this file segment
		for leftOffset := 0; leftOffset < len(newDiskSegments)-rightOffset; leftOffset++ {
			if newDiskSegments[leftOffset].fileID == -1 && newDiskSegments[leftOffset].size >= currentFileSegment.size {

				freeSpaceLeftOver := newDiskSegments[leftOffset].size - currentFileSegment.size

				// move the file segment to the left side
				newDiskSegments[leftOffset].fileID = currentFileSegment.fileID
				newDiskSegments[leftOffset].size = currentFileSegment.size

				// mark the right side as freeSpace segment
				newDiskSegments[len(newDiskSegments)-rightOffset].fileID = -1

				// if there are still space left, add it as new Segment to the 'newDiskSegments'
				if freeSpaceLeftOver > 0 {
					newDiskSegments = slices.Insert(newDiskSegments, leftOffset+1, DiskSegment{-1, freeSpaceLeftOver})
				}
				break
			}
		}
		rightOffset += 1
	}

	return newDiskSegments
}

func main() {
	rawDisk := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		rawDisk += line
	}

	fmt.Println("original input: ", rawDisk)

	// parse the raw disk format into a linked list where each node is either file (id= fileID, size=fileSize) or free space (id = -1, size=freespaceSize)
	diskBlocks := ParseRawDisk(rawDisk)

	fmt.Printf("diskBlocks before defragmentation: %v\n", diskBlocks)

	newDiskBlocks := DefragmentWithWholeFileMove(diskBlocks)

	fmt.Printf("diskBlocks after defragmentation: %v\n", newDiskBlocks)

	sliceOfUnitDiskBlocks := make([]int, 0)
	for _, segment := range newDiskBlocks {
		if segment.fileID == -1 {
			for i := 0; i < segment.size; i++ {
				sliceOfUnitDiskBlocks = append(sliceOfUnitDiskBlocks, -1)
			}
		} else {
			for i := 0; i < segment.size; i++ {
				sliceOfUnitDiskBlocks = append(sliceOfUnitDiskBlocks, segment.fileID)
			}
		}
	}

	// calculate checksum
	var checksum uint64 = 0
	for i, fileID := range sliceOfUnitDiskBlocks {
		if fileID == -1 {
			continue
		}
		checksum += uint64(i * fileID)
	}
	// print checksum
	fmt.Println("checksum: ", checksum)
}
