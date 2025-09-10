package main

import "fmt"

func main() {
	s := Solution{}
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	result := s.Rotate(nums, k)
	fmt.Println(result) // Output: [5,6,7,1,2,3,4]

	// Generate Pascal's Triangle with 5 rows
	v := _Solution{}
	resultTriangle := v.Generate(5)
	fmt.Println(resultTriangle)
}

// Solution is a struct to define the receiver type
type Solution struct{}

func (Solution) Rotate(nums []int, k int) []int {
	reverseSection(nums, 0, len(nums)-1)
	reverseSection(nums, 0, k-1)
	reverseSection(nums, k, len(nums)-1)
	return nums
}

func reverseSection(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start] // swap
		start++
		end--
	}
}
