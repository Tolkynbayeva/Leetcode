package main

func isPerfectSquare(num int) bool {

	if num < 0 {
		return false
	} else if num == 0 || num == 1 {
		return true
	}

	for i := 0; i < num; i++ {
		if (i * i) == num {
			return true
			break
		}
	}
	return false
}

// func main() {
// 	fmt.Println(isPerfectSquare(16)) // true
// 	fmt.Println(isPerfectSquare(14)) // false
// 	fmt.Println(isPerfectSquare(1))  // true
// }
