package main

import (
	"fmt"
	"math/rand"
)

/*
*	扇動の定義：隣り合う要素の入れ替え
*	終了条件：暫定解の更新回数はN-1回(N:要素数)とする
*/

var job_time = [4][4]int{{6, 1, 9, 3}, {2, 5, 7, 8}, {6, 3, 5, 4}, {3, 5, 2, 1}}
var p int = 5

func shuffle(data []int) []int{
    for l := len(data) - 1; l >= 0; l-- {
        j := rand.Intn(l + 1)
        data[l], data[j] = data[j], data[l]
	}
	return data
}

func swap(c []int, i int) []int{
	swap_c := c
	swap_c[i], swap_c[i+1] =  swap_c[i+1], swap_c[i]
	return swap_c
}

func mslsearch(comb []int) (int, []int){
	first_value := job_time[0][comb[0]-1] + job_time[1][comb[1]-1] + job_time[2][comb[2]-1] + job_time[3][comb[3]-1]
	best_c := make([]int, len(comb))

	result_value := first_value
	for i := 0; i < len(comb) - 1; i++ {
		swap_comb := swap(comb, i)
		new_value := job_time[0][swap_comb[0]-1] + job_time[1][swap_comb[1]-1] + job_time[2][swap_comb[2]-1] + job_time[3][swap_comb[3]-1]
		if result_value > new_value {
			result_value = new_value
			_ = copy(best_c, swap_comb)
		}
	}
	return result_value, best_c
}

func main(){
	comb := []int{3, 4, 1, 2}
	shuffle_comb := []int{}
	optimal := []int{}
	best_comb := []int{}
	optimal_value := 0

	for k := 0; k < p; k++{
		shuffle_comb = shuffle(comb)
		result, best_c := mslsearch(shuffle_comb)
		optimal = append(optimal, result)
		
		if k == 0{
			optimal_value = optimal[0]
			best_comb = best_c
		}else if k < 0 && optimal[k-1] > optimal[k]{
			optimal_value = optimal[k]
			best_comb = best_c
		}
	}

	fmt.Println("準暫定解", best_comb)
	fmt.Println("目的関数値", optimal_value)
}