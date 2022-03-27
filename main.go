package main

import (
	"fmt"
)

func mapSlice[T any, U any](slice []T, mapper func(item T) U) []U {
	newSlice := []U{}
	for _, item := range slice {
		newSlice = append(newSlice, mapper(item))
	}
	return newSlice
}

func main() {

	list := []int{1, 2, 3}
	newSl := mapSlice(list, func(item int) int {
		return item * 2
	})

	for _, item := range newSl {
		fmt.Printf("item: %v\n", item)
	}

}
