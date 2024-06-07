package main

import "fmt"

func transpose(matrix [][]int) {
	for i := range matrix {
		for j := i + 1; j < len(matrix[i]); j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = t
		}
	}
}
func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, v := range mat {
		fmt.Println(v)
	}
	fmt.Println("-----")
	transpose(mat)
	for _, v := range mat {
		fmt.Println(v)
	}
}
