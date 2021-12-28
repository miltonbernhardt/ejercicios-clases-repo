package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	isFinishSortByInsertion := make(chan bool)
	go sortByInsertion(variable1, variable2, variable3, isFinishSortByInsertion)
	isFinishSortBySelection := make(chan bool)
	go sortBySelection(variable1, variable2, variable3, isFinishSortBySelection)


	isFinishSortByBubble := make(chan bool)
	go sortByBubble(variable1, variable2, variable3, isFinishSortByBubble)

	<-isFinishSortBySelection
	<-isFinishSortByInsertion
	<-isFinishSortByBubble
}

func sortByInsertion(variable1 []int, variable2 []int, variable3 []int, isFinish chan bool) {
	start := time.Now()
	InsertionSort(variable1)
	InsertionSort(variable2)
	InsertionSort(variable3)
	fmt.Printf("INSERTION SORT TIME: %v\n", time.Since(start))
	isFinish <- true
}

func sortByBubble(variable1 []int, variable2 []int, variable3 []int, isFinish chan bool) {
	start := time.Now()
	BubbleSort(variable1)
	BubbleSort(variable2)
	BubbleSort(variable3)
	fmt.Printf("BUBLE SORT TIME:  = %v\n", time.Since(start))
	isFinish <- true
}

func sortBySelection(variable1 []int, variable2 []int, variable3 []int, isFinish chan bool) {
	start := time.Now()
	SelectionSort(variable1)
	SelectionSort(variable2)
	SelectionSort(variable3)
	fmt.Printf("SELECTION SORT TIME:  %v\n", time.Since(start))
	isFinish <- true
}

func InsertionSort(slice []int) {
	var i, j int
	for i = 0; i < len(slice); i++ {
		actual := slice[i]
		for j = i; j > 0 && slice[j-1] > actual; j-- {
			slice[j] = slice[j-1]
		}
		slice[j] = actual
	}
}

func BubbleSort(slice []int) {
	iteration := 0
	var changesMade = true
	for changesMade {
		changesMade = false
		iteration++
		for j := 0; j < len(slice)-iteration; j++ {
			if slice[j] > slice[j+1] {
				actual := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = actual
				changesMade = true
			}
		}
	}
}

func SelectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
}
