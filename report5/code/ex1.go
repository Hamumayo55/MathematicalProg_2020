package main

import (
	"fmt"
	"math"
)
var num_row int = 3
var num_line int = 4
var max_z float64 = 0
var pe float64 = 0
var matrix = [][]float64{{8,1,-1,2},{1,2,-3,-1},{0,-1,-1,-1}}

func peselect(m [][]float64){
	row := 0
	line := 0

	// 列の選択
	for i := 1; i < num_row+1; i++{
		if i != 1{
			if math.Abs(m[2][i]) > max_z{
				row = i
			}
		}else{
			row = i
			max_z = math.Abs(m[2][i])
		}
	}
	// ピボットエレメントの選択
	for n := 0; n < num_row-1; n++{
		if n != 0{
			if pe > (m[n][0]/m[n][row]){
				pe = m[n][0]/m[n][row]
				line = n
			}
		}else{
			pe = m[n][0]/m[n][row]
			line = n
		}
	}
	fmt.Println(line)
	fmt.Println(pe)
}

func main(){
	peselect(matrix)
}