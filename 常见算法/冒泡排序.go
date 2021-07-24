package main

import "fmt"

func main()  {
	var arr = [...]int{99, 51, 41, 2, 31}

	var rarr = bubbleSort(arr[:])
	fmt.Println(rarr)
}

func bubbleSort(slice []int) []int {
	for n := 0; n <= len(slice); n++ {
		for i := 1; i < len(slice)-n; i++ {
			if slice[i] < slice[i-1] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
			}
		}
	}
	return slice
}