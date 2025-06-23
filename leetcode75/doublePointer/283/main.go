package main

import "fmt"

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

}
