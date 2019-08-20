package GoLeetcode

//70 爬楼梯

func climbStairs(n int) int {
	a := 1
	b := 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b

}

//120 三角形最小路径和
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= len(triangle[row])-1; col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}
