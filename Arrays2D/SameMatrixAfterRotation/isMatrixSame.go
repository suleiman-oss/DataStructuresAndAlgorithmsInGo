package main

import "fmt"

//https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/description/
func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		RotateClockwise(mat)
		if isSame(mat, target) {
			return true
		}
	}
	return false
}

func isSame(mat, target [][]int) bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if i >= len(target) || j >= len(target[i]) {
				return false
			}
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}

func RotateClockwise(matrix [][]int) {
	transposeOfSquareMat(matrix)
	for i := range matrix {
		s, e := 0, len(matrix[i])-1
		for s < e {
			t := matrix[i][s]
			matrix[i][s] = matrix[i][e]
			matrix[i][e] = t
			s++
			e--
		}
	}
}

func transposeOfSquareMat(matrix [][]int) {
	for i := range matrix {
		for j := i + 1; j < len(matrix[i]); j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = t
		}
	}
}

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	t := [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(findRotation(mat, t))
}
