package main

import (
	"fmt"
	//"sort"
	"github.com/voidlotus/mysearch"
	"github.com/voidlotus/mysort"
)

func printArray(arr []int, comments string) {
	fmt.Println(comments)
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

func main() {
	//testCase1 := []int{1, 4, 3, 2, 5}
	//sort.Sort(sort.IntSlice(testCase1)) // cara 2
	//sort.Ints(testCase1) // cara 1
	//res := binarysearch.Search(testCase1, 4)
	//res := 1
	//fmt.Printf("\nResult found at index: %d", res)

	testCase2 := []int{3, 0, 4, 2, 1, 5}
	printArray(testCase2, "Before merge sorting: ")
	mysort.MergeSort(testCase2, 0, len(testCase2)-1)
	printArray(testCase2, "After merge sorting: ")

	testCase3 := []int{3, 0, 4, 2, 1, 5}
	printArray(testCase3, "Before quick sorting: ")
	mysort.MergeSort(testCase3, 0, len(testCase3)-1)
	printArray(testCase3, "After quick sorting: ")

	testCase4 := [][]int{{1, 2}, {0, 2, 3}, {0, 1, 3, 4}, {1, 2, 4}, {2, 3}}
	visited := make([]bool, (len(testCase4)))
	mysearch.DFS(testCase4, 0, visited)

	fmt.Println("\nBFS Traversal: ")
	testCase5 := [][]int{{1, 2}, {0, 2, 3}, {0, 1, 3, 4}, {1, 2, 4}, {2, 3}}
	visited2 := make([]bool, (len(testCase5)))
	mysearch.BFS(testCase5, 0, visited2)
}
