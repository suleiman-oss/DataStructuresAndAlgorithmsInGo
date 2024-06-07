package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	ans := make([][]int, len(matrix[0]))
	for i := range ans {
		ans[i] = make([]int, len(matrix))
	}
	for i := range ans {
		for j := range ans[i] {
			ans[i][j] = matrix[j][i]
		}
	}
	return ans
}
func transposeOfSquareMat(matrix [][]int) [][]int {
	for i := range matrix {
		for j := i + 1; j < len(matrix[i]); j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = t
		}
	}
	return matrix
}
func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, v := range mat {
		fmt.Println(v)
	}
	fmt.Println("-----")
	m := transposeOfSquareMat(mat)
	for _, v := range m {
		fmt.Println(v)
	}
}
