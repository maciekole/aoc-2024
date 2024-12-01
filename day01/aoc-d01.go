package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func loadData() ([]int, []int, error) {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		return nil, nil, err
	}

	r := bufio.NewReader(file)

	var listA []int
	var listB []int

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			items := strings.Split(string(line[:]), "   ")
			itemA, errA := strconv.Atoi(items[0])
			if errA != nil {
				return nil, nil, errA
			}
			itemB, errB := strconv.Atoi(items[1])
			if errB != nil {
				return nil, nil, errB
			}
			listA = append(listA, itemA)
			listB = append(listB, itemB)
		}
		if err != nil {
			break
		}
	}

	return listA, listB, nil
}

func count(item int, slice []int) int {
	count := 0
	for _, s := range slice {
		if s == item {
			count++
		}
	}
	return count
}
func main() {
	totalDistance := 0
	similarity := 0
	similarityMap := make(map[int]int)

	listA, listB, err := loadData()
	if err != nil {
		fmt.Println("Input file 'input.txt' not found!")
	}

	sort.Sort(sort.IntSlice(listA))
	sort.Sort(sort.IntSlice(listB))

	if len(listA) == len(listB) {
		for i := range listA {
			// Part 1
			if listA[i] >= listB[i] {
				totalDistance = totalDistance + (listA[i] - listB[i])
			} else {
				totalDistance = totalDistance + (listB[i] - listA[i])
			}
			sim, ok := similarityMap[listA[i]]
			fmt.Printf("listA[i]: %v\n", listA[i])
			if ok {
				fmt.Printf("sim: %v\n", totalDistance)
				similarity = similarity + sim
			} else {
				iCount := count(listA[i], listB)
				if iCount != 0 {
					iSim := listA[i] * iCount
					fmt.Printf("iCount: %v\n", iCount)
					fmt.Printf("iSim: %v\n", iSim)
					similarityMap[listA[i]] = iSim
					similarity = similarity + iSim
				}
			}

		}
	}
	fmt.Printf("Total distance: %v\n", totalDistance)
	fmt.Printf("Similarity: %v\n", similarity)
}
