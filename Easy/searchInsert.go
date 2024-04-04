package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	return l
}

// func main() {
//     fmt.Println(searchInsert([]int{1,3,5,6}, 5)) // 2
//     fmt.Println(searchInsert([]int{1,3,5,6}, 2)) // 1
//     fmt.Println(searchInsert([]int{1,3,5,6}, 7)) // 4
// }
