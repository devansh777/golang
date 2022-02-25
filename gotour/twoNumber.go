package main

import "fmt"

func main() {
	sumRange := []int{2, 3, 1, 4, 2}
	fmt.Println(twoSum(sumRange, 5))
}
func twoSum(nums []int, target int) []int {
	sum := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sum = append(sum, i)
				sum = append(sum, j)
				break
			}
		}
	}
	return sum
}
