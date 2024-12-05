// https://adventofcode.com/2024/day/5
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Page struct {
	num          int
	predecessors []int // the pages that needs to be printed before this page, i.e this page's depedencies
}

// convert a slice of strings to a slice of integers
func ConvertToInts(pages []string) []int {
	ints := make([]int, len(pages))
	for i, page := range pages {
		n, err := strconv.Atoi(page)
		if err != nil {
			panic(err)
		} else {
			ints[i] = n
		}
	}
	return ints
}

func IsLegalPageUpdate(pages []int, pageRule map[int]Page) bool {
	// check if the page update is valid
	// for each page in the update, check if any pages after it are in its predecessors
	for i, pageNum := range pages {
		if currentPage, ok := pageRule[pageNum]; ok {
			// current page has predecessors in the pageRule, let's check all the pages after it
			for _, pageAfter := range pages[i+1:] {
				if slices.Contains(currentPage.predecessors, pageAfter) {
					// pageAfter is a predecessor of pageNum, which is not allowed
					return false
				}
			}
		}
	}

	return true
}

func SortPageByRule(pages []int, pageRule map[int]Page) []int {
	// Use bubble sort to sort the pages
	for i := 0; i < len(pages)-1; i++ {
		swapped := false

		for j := 0; j < len(pages)-i-1; j++ {
			// if page[j+1] is in predecessors of page[j], swap the pages
			if slices.Contains(pageRule[pages[j]].predecessors, pages[j+1]) {
				// swap the pages
				pages[j], pages[j+1] = pages[j+1], pages[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return pages
}

func main() {
	pageOrders := make(map[int]Page)

	// read the page order in 'page_num_1|page_num_2' form
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scan input line by line
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}

		var pre, post int
		fmt.Sscanf(line, "%d|%d", &pre, &post)

		if existingPage, ok := pageOrders[post]; !ok {
			// no such page exists, create a new page
			pageOrders[post] = Page{num: post, predecessors: []int{pre}}
		} else {
			// update existing page
			existingPage.predecessors = append(existingPage.predecessors, pre)
			pageOrders[post] = existingPage
		}
	}

	// // print the page order rules
	// for pageNum, page := range pageOrders {
	// 	fmt.Printf("Page %d: %v\n", pageNum, page.predecessors)
	// }

	// now read in the page updates, each update is in the form of 'page_1,page_2,...,page_n'
	pageUpdates := make([][]int, 0)
	for scanner.Scan() {
		// scan input line by line
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		// split the line by comma
		pageNumbers := ConvertToInts(strings.Split(line, ","))
		pageUpdates = append(pageUpdates, pageNumbers)
	}

	fmt.Println("Number of Page updates: ", len(pageUpdates))

	validPageUpdates := make([][]int, 0)
	invalidPageUpdates := make([][]int, 0)

	for _, pageUpdate := range pageUpdates {
		// check page update against the page order,
		validUpdate := IsLegalPageUpdate(pageUpdate, pageOrders)
		fmt.Println("Page update: ", pageUpdate, "valid: ", validUpdate)

		if validUpdate {
			validPageUpdates = append(validPageUpdates, pageUpdate)
		} else {
			invalidPageUpdates = append(invalidPageUpdates, pageUpdate)
		}
	}

	middlePageSum := 0
	for _, pageUpdate := range validPageUpdates {
		middlePageSum += pageUpdate[len(pageUpdate)/2]
	}
	// print the middle page sum
	fmt.Println("Middle page sum of validPageUpdates: ", middlePageSum)

	middlePageSumOfCorrections := 0
	for _, pageUpdate := range invalidPageUpdates {
		// sort the page update based on the page order
		sortedPageUpdate := SortPageByRule(pageUpdate, pageOrders)
		middlePageSumOfCorrections += sortedPageUpdate[len(sortedPageUpdate)/2]
	}

	fmt.Println("Middle page sum of corrections: ", middlePageSumOfCorrections)

}
