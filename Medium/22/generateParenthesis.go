func generateParenthesis(n int) []string {
	result := []string{}
	back(&result, "", 0, 0, n)
	return result
}

func back(result *[]string, curr string, open int, close int, n int) {
	if len(curr) == n*2 {
		*result = append(*result, curr)
		return
	}

	if open < n {
		back(result, curr+"(", open+1, close, n)
	}
	if close < open {
		back(result, curr+")", open, close+1, n)
	}
}