package main

import (
	"fmt"
)

/*
*	扇動の定義：隣り合う要素の入れ替え
*	終了条件：隣り合う同士がすべて入れ替わったら
*/

var job_time = [4][4]int{{6, 1, 9, 3}, {2, 5, 7, 8}, {6, 3, 5, 4}, {3, 5, 2, 1}}

func swap(c []int, i int) []int{
	swap_c := make([]int, len(c))
	_ = copy(swap_c, c)
	swap_c[i], swap_c[i+1] =  swap_c[i+1], swap_c[i]
	return swap_c
}

func mslsearch(comb []int) (int, int){
	first_value := job_time[0][comb[0]-1] + job_time[1][comb[1]-1] + job_time[2][comb[2]-1] + job_time[3][comb[3]-1]

	result_value := first_value
	for i := 0; i < len(comb) - 1; i++ {
		swap_comb := swap(comb, i)
		new_value := job_time[0][swap_comb[0] - 1] + job_time[1][swap_comb[1] - 1] + job_time[2][swap_comb[2] - 1] + job_time[3][swap_comb[3] - 1]
		if result_value > new_value {
			result_value = new_value
		}
	}
	return first_value, result_value
}

func main(){

	comb := []int{3, 4, 1, 2}
	first, result := mslsearch(comb)
	fmt.Println("初期解", first)
	fmt.Println("最良解", result)
}