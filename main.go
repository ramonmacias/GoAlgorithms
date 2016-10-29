package main

import (
	"log"

	"github.com/GoAlgorithms/sorting"
)

func main() {
	var s []int
	s = append(s, 11)
	s = append(s, 1)
	s = append(s, 45)
	s = append(s, 66)
	s = append(s, 79978)
	s = append(s, 3)
	s = append(s, 8)
	s = append(s, 68)
	s = append(s, 90)
	log.Printf("Unordered list: %v", s)
	orderedSlice := sorting.IntSort(s)
	log.Printf("Ordered list using insertion sort method: %v", orderedSlice)
}
