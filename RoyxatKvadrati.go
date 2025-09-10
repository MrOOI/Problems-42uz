package main

type RoyxatKvadrati struct{}

func (_ RoyxatKvadrati) SortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = nums[i] * nums[i]
		nums[i] = result[i]
	}
	start := 0
	end := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[start] > nums[end] {
			result[i] = nums[start]
			start++
		} else {
			result[i] = nums[end]
			end--
		}
	}

	return result
}
