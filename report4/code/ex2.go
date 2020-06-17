package main

import (
	"fmt"
	"math/rand"
)

/*
*	Definition of perturbation : replacement of adjacent elements
*	End condition : Tentative solution updated N-1 times (N:number of elements)
*/

var job_time = [4][4]int{{6, 1, 9, 3}, {2, 5, 7, 8}, {6, 3, 5, 4}, {3, 5, 2, 1}} 
var first_comb = []int{} //Objective function value of the first solution
var optimal_comb = []int{} //Objective function value of the end solution
var p int = 5 //The number of first solutions to be generated
var memory_min_first = make([]int, len(job_time[0])) //The best of the first solution
var memory_max_first = make([]int, len(job_time[0])) //The worst of the first solution
//Minimum and maximum judgment
var min_first int = 1000 
var max_first int = 0
var first_flag bool = true

//sort function
func sort(s []int) []int {
	for i := 0; i < len(s) - 1; i++ {
		for j := 0; j < len(s) - i - 1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

//shuffle function
func shuffle(data []int) []int{
    for l := len(data) - 1; l >= 0; l-- {
        j := rand.Intn(l + 1)
        data[l], data[j] = data[j], data[l]
	}
	return data
}

//perturbation function
func swap(c []int, i int) []int{
	swap_c := c
	swap_c[i], swap_c[i+1] =  swap_c[i+1], swap_c[i]
	return swap_c
}

//judge the best and worst of the first solutions
func memory_slice(f int, c []int){
	if first_flag {
		min_first, max_first = f, f
		first_flag = false
	}else if !first_flag{
		if min_first > f {
			min_first = f
			_ = copy(memory_min_first, c)
		}else if max_first < f{
			max_first = f
			_ = copy(memory_max_first, c)
		}
	}
}

//Multi Start local search
func mslsearch(comb []int) (int, []int){
	first_value := job_time[0][comb[0]-1] + job_time[1][comb[1]-1] + job_time[2][comb[2]-1] + job_time[3][comb[3]-1]

	memory_slice(first_value, comb)

	best_c := make([]int, len(comb))
	_ = copy(best_c, comb)

	new_value := 0

	result_value := first_value
	for i := 0; i < len(comb) - 1; i++ {
		swap_comb := swap(comb, i)
		new_value = job_time[0][swap_comb[0]-1] + job_time[1][swap_comb[1]-1] + job_time[2][swap_comb[2]-1] + job_time[3][swap_comb[3]-1]
		if result_value > new_value {
			result_value = new_value
			_ = copy(best_c, swap_comb)
		}
	}
	return result_value, best_c
}

func main(){
	comb := []int{3, 4, 1, 2} //first comb
	shuffle_comb := []int{} //shuffle
	best_comb := make([]int, len(job_time[0]))
	bad_comb := make([]int, len(job_time[0]))

	for k := 0; k < p; k++{
		shuffle_comb = shuffle(comb)
		result, best_c := mslsearch(shuffle_comb)
		optimal_comb = append(optimal_comb, result)

		if k == 0{
			_ = copy(best_comb, best_c)
		}else if k > 0 {
			if optimal_comb[k-1] > optimal_comb[k]{
				_ = copy(best_comb, best_c)
			}else if optimal_comb[k-1] < optimal_comb[k]{
				_ = copy(bad_comb, best_c)
			}
		}
		optimal_comb = sort(optimal_comb)
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println("初期解の中での最良な解", memory_min_first, "|", "目的関数値", min_first)
	fmt.Println("初期解の中での最悪な解", memory_max_first, "|", "目的関数値", max_first)
	fmt.Println("--------------------------------------------------")
	fmt.Println("最終解の中での最良な解", best_comb, "|", "目的関数値", optimal_comb[0])
	fmt.Println("最終解の中での最悪な解", bad_comb, "|", "目的関数値", optimal_comb[len(optimal_comb)-1])
	fmt.Println("--------------------------------------------------")
}