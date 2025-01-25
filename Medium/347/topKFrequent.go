package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(nums); i >= 0 && len(result) < k; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}
	return result[:k]
}

func main() {
	nums := []int{1,1,1,2,2,3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}