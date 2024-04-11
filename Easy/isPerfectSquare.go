package main

func isPerfectSquare(num int) bool {
	l := 1
	r := num
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// func main() {
// 	fmt.Println(isPerfectSquare(16)) // true
// 	fmt.Println(isPerfectSquare(14)) // false
// 	fmt.Println(isPerfectSquare(1))  // true
// }
