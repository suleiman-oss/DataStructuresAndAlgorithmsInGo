package main

import "fmt"

//https://leetcode.com/problems/rotate-image/description/

//Transpose and reverse each row
func RotateClockwise(matrix [][]int) [][]int {
	transpose := transposeOfSquareMat(matrix)
	for i := range transpose {
		s, e := 0, len(transpose[i])-1
		for s < e {
			t := transpose[i][s]
			transpose[i][s] = transpose[i][e]
			transpose[i][e] = t
			s++
			e--
		}
	}
	return transpose
}

//Transpose and reverse each column
func RotateCounterClockwise(matrix [][]int) [][]int {
	transpose := transposeOfSquareMat(matrix)
	ans := make([][]int, len(transpose))
	for i := range ans {
		ans[i] = make([]int, len(transpose[i]))
	}
	for i := 0; i < len(transpose); i++ {
		ans[len(ans)-1-i] = transpose[i]

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
	m := RotateCounterClockwise(mat)
	for _, v := range m {
		fmt.Println(v)
	}
}
