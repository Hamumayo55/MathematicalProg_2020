package main

import (
	"fmt"
)

/*
*	扇動の定義：隣り合う要素の入れ替え
*	終了条件：隣り合う同士がすべて入れ替わったら
*/

var job_time = [4][4]int{{6, 1, 9, 3}, {2, 5, 7, 8}, {6, 3, 5, 4}, {3, 5, 2, 1}}

func mslsearch(c []int) int{
	first_value := job_time[0][c[0] - 1] + job_time[1][c[1] - 1] + job_time[2][c[2] - 1] + job_time[3][c[3] - 1]
	return kotae
}

func main(){

	comb := []int{3, 4, 1, 2}
	result := mslsearch(comb)
	fmt.Println(result)
}