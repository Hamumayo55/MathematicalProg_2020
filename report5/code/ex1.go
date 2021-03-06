package main

import (
	"fmt"
	"math"
)
var num_row int = 4 //列
var num_line int = 3 //行
var max_z float64 = 0
var pe_div float64 = 0
var pe float64 = 0
var matrix = [][]float64{{8,1,-1,2},{1,2,-3,-1},{0,-1,-1,-1}}

func simplex(row int, line int) [][]float64{
	for i := 0; i < num_row; i++{
		matrix[line][i] = matrix[line][i]/pe
	}
	for n := 0; n < num_line; n++{
		p_line := matrix[n][row]
		for j := 0; j < num_row; j++{
			if n != line{
				matrix[n][j] = matrix[n][j]-matrix[line][j]*p_line
			}
		}
	}
	fmt.Println("X",row,"⇆ X",line+4)
	return matrix
}

func peselect(m [][]float64)(int, int, bool){
	row := 0
	line := 0
	flag := false

	// 列の選択
	for i := 1; i < num_row; i++{
		if i != 1{
			if math.Abs(m[num_line-1][i]) > max_z{
				fmt.Println()
				row = i
				max_z = math.Abs(m[num_line-1][i])
			}
		}else{
			row = i
			max_z = math.Abs(m[num_line-1][i])
		}
	}
	for n := 0; n < num_line; n++{
		if matrix[n][row] == 0 || matrix[n][row] < 0{
			flag = true
		}else{
			flag = false
			break
		}
	}
	// ピボットエレメントの選択
	if !flag{
		for n := 0; n < num_line-1; n++{
			if n != 0 && m[n][0]/m[n][row] > 0{
				if pe_div > (m[n][0]/m[n][row]){
					pe_div = m[n][0]/m[n][row]
					line = n
					pe = m[n][row]
				}
			}else if m[n][0]/m[n][row] > 0 {
				pe_div = m[n][0]/m[n][row]
				line = n
				pe = m[n][row]
			}
		}
	}
	return row, line, flag
}

func main(){
	fmt.Println("初期のシンプレックスタブロー")
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])
	fmt.Println("---------------------------------------------------")

	z_flag := false
	isoptimal := true

	for{
		row, line, flag := peselect(matrix)
		if flag {
			isoptimal = false
			break
		}
		matrix = simplex(row, line)
		fmt.Println("シンプレックスタブロー")
		fmt.Println(matrix[0])
		fmt.Println(matrix[1])
		fmt.Println(matrix[2])
		fmt.Println("---------------------------------------------------")

		for i := 0; i < num_row; i++{
			if matrix[num_line-1][i] == 0 || matrix[num_line-1][i] > 0{
				z_flag = true
			}else{
				z_flag = false
				break
			}
		}
		if z_flag {
			break
		}
	}
	if isoptimal{
		fmt.Println("最終シンプレックスタブロー")
		fmt.Println(matrix[0])
		fmt.Println(matrix[1])
		fmt.Println(matrix[2])
		fmt.Println("---------------------------------------------------")
	}else{
		fmt.Println("最終シンプレックスタブロー")
		fmt.Println(matrix[0])
		fmt.Println(matrix[1])
		fmt.Println(matrix[2])
		fmt.Println("---------------------------------------------------")
		fmt.Println("最適化なし")
	}
}