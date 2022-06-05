package main

import "fmt"

func main() {
	fmt.Println("Start rotating array")
	matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Printf("matrix is :%v\n", matrix)
	rotate(matrix)
}

func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	fmt.Println(n)
	var rotatedMatrix [][]int
	for i := 0; i < n; i++ {
		row := make([]int, n)
		rotatedMatrix = append(rotatedMatrix, row)
	}
	fmt.Println(rotatedMatrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row := j
			col := n - 1 - i
			rotatedMatrix[row][col] = matrix[i][j]
		}
	}
	fmt.Println(rotatedMatrix)
	return nil
}
