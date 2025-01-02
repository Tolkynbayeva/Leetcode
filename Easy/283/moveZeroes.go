package main

func moveZeroes(nums []int) {
	l := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
