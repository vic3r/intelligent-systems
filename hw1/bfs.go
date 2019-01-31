package main

import (
	"fmt"
	"reflect"

	"github.com/golang-collections/collections/set"

	"github.com/golang-collections/collections/queue"
)

func breadFirstSearch(initialState, goalTest []int) bool {

	frontier := queue.New()
	frontier.Enqueue(initialState)
	explored := set.New()
	visitedNodes := 0
	depth := 0

	fmt.Println("Path to goal: ")
	for frontier.Len() > 0 {
		state := frontier.Dequeue().([]int)
		fmt.Println(state)
		explored.Insert(state)

		if reflect.DeepEqual(initialState, goalTest) {
			fmt.Printf("Number of visited nodes %d", visitedNodes)
			fmt.Printf("The cost of the path is %d", depth)
			return true
		}

		for neighbor := range state {
			if !explored.Has(neighbor) {
				frontier.Enqueue(neighbor)
				explored.Insert(neighbor)
				depth++
			}
		}
		visitedNodes++
	}

	return false
}

func main() {
	initialState := []int{1, 2, 5, 3, 4, 0, 6, 7, 8}
	goalTest := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(breadFirstSearch(initialState, goalTest))
}
