package mysearch

import "fmt"

func BinarySearch(arr []int, target int) (res int) {
	low := 0
	high := len(arr) - 1
	//fmt.Println("Searching...")
	//fmt.Printf("low/high: %d/%d (target: %d)\n", low, high, target)
	for low <= high {
		mid := low + (high-low)/2
		//fmt.Printf("%d ", arr[mid])
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func DFS(arr [][]int, currentVertex int, visited []bool) {
	
	visited[currentVertex] = true
	fmt.Printf("%d ", currentVertex)
	for _,neighbor := range arr[currentVertex] {
		if !visited[neighbor] {
			DFS(arr, neighbor, visited)
		}
	} 
}

func BFS(arr [][]int, startVertex int, visited []bool){
	visited[startVertex] = true
	// create queue using Channel
	ch := make(chan int, len(arr))
	ch <- startVertex
	v, ok := <-ch
	for !ok {
		currentVertex := v

		fmt.Printf("%d ", v)

		for neighbor := range arr[currentVertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				ch <- neighbor
			}
		}

	}
}
