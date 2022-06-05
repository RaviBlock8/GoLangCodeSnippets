package main

/**
* 1 2 3
* 4 5 6
* 7 8 9
**/

import "fmt"

func main() {
	fmt.Println("Starting....")
	matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
}

func spiralOrder(matrix [][]int) []int {
	forward := true
	i := 0
	j := 0
	col := len(matrix[0])
	up := 0
	down := len(matrix) - 1
	for true {
		if forward && i == down {
			forward = false
			up += 1
		}
		if !forward && i == up {
			forward = true
			down -= 1
		}
		if forward && i == up {
			for _, e := range matrix[up] {
				fmt.Println(e)
			}
			// matrix = append(matrix[1:])
		} else if !forward && i == down {
			for _, e := range matrix[i] {
				fmt.Println(e)
			}
			// matrix = append(matrix[:j])
		} else {
			fmt.Println(matrix[i][j])
			if j == 0 {
				matrix[i] = append(matrix[i][1:])
			}
		}

		if forward {
			i++
			j = col
		} else {
			i--
			j = 0
		}
	}
}
