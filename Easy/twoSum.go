package main

func twoSum(nums []int, target int) []int {
	var result []int
	m := make(map[int]int)

	for i, n := range nums {
		f := target - n
		if j, ok := m[f]; ok {
			return []int{j, i}
		}
		m[n] = i
	}

	return result
}

// func main() {
// 	fmt.Println(twoSum([]int{2,7,11,15}, 9)) // [0,1]
// 	fmt.Println(twoSum([]int{3,2,4}, 6)) // [1,2]
// 	fmt.Println(twoSum([]int{3,3}, 6)) // [[0,1]
// }
