package main

func getConcatenation(nums []int) []int {
	result := nums
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i])
	}
	return result
}

// func main() {
// 	fmt.Println(getConcatenation([]int{1, 2, 1})) // Output: [1,2,1,1,2,1]
// 	fmt.Println(getConcatenation([]int{1,3,2,1})) // Output: [1,3,2,1,1,3,2,1]
// }