package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	for _, count := range m {
		if count > 1 {
			return true
		}
	}
	return false
}

// func main() {
// 	fmt.Println(containsDuplicate([]int{1, 2, 3, 1})) // true
// 	fmt.Println(containsDuplicate([]int{1, 2, 3, 4})) // false
// 	fmt.Println(containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2})) // true
// }
