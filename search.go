package main

import "fmt"

func linearSearch(arr []int, target int) int {
	i := 0
	for i = 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
func main() {
	//arr  := {12,23}
	fmt.Println("PUSH TEST")
	fmt.Println("PUSH TEST")
	fmt.Println("PUSH TEST")
	fmt.Println("PUSH TEST")

}
