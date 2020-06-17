package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*	Definition of perturbation : replacement of adjacent elements
*	End condition : Tentative solution updated N-1 times (N:number of elements)
*	Choose the first solution : The solution obtained by the greedy method is the first solution
*/

var size = []int{3,6,5,4,8,5,3,4,3,5,6,4,8,7,11,8,14,6,12,4}
var price = []int{7,12,9,7,13,8,4,5,3,10,7,5,6,14,5,9,6,12,5,9}
//Minimum and maximum judgment
var min_first_p int = 1000 
var max_first_p int = 0
var min_first_s int = 0 
var max_first_s int = 0
var memory_min_first = make([]int, len(size)) //The best of the first solution
var memory_max_first = make([]int, len(size)) //The worst of the first solution
var first_flag bool = true
var g_max_size int = 0
var g_max_price int = 0
var limit int =  55

//sort function
func sort(s []int, p []int)([]int, []int){
	for i := 0; i < len(s) - 1; i++ {
		for j := 0; j < len(s) - i - 1; j++ {
			if float64(p[j])/float64(s[j]) < float64(p[j + 1])/float64(s[j + 1]) {
				s[j], s[j + 1] = s[j + 1], s[j]
				p[j], p[j + 1] = p[j + 1], p[j]
			}
		}
	}
	return s, p
}

//shuffle function
func shuffle(data []int) []int{
	rand.Seed(time.Now().Unix())
	n := len(data)
    for i := n-1; i >= 0; i-- {
		j := rand.Intn(i+1)
		data[i], data[j] = data[j], data[i]
	}
	return data
}

//greedy function
func greedy(s []int, p []int)[]int{
	g_comb := []int{}
	for n := 0; n < len(s) - 1; n++{
		g_max_size += s[n]
		g_max_price += p[n]
		if g_max_size > limit{
			g_max_size -= s[n]
			g_max_price -= p[n]
			g_comb = append(g_comb, 0)
		}else{
			g_comb = append(g_comb, 1)
		}
	}
	return g_comb
}


//perturbation function
func swap(c []int, i int) []int{
	swap_c := c
	swap_c[i], swap_c[i+1] =  swap_c[i+1], swap_c[i]
	return swap_c
}

//judge the best and worst of the first solutions
func memory_slice(f_p int, f_s int,c []int){ 
	if first_flag {
		min_first_p, max_first_p = f_p, f_p
		min_first_s, max_first_s = f_s, f_s
		_ = copy(memory_max_first, c)
		_ = copy(memory_min_first, c)
		first_flag = false
	}else if !first_flag{
		if min_first_p > f_p {
			min_first_p = f_p
			min_first_s = f_s
			_ = copy(memory_min_first, c)
		}else if max_first_p < f_p{
			max_first_p = f_p
			max_first_s = f_s
			_ = copy(memory_max_first, c)
		}	
	}
}

//Multi Start local search
func mslsearch(c []int, limit int) (int, int, []int){
	max_size := 0
	max_price := 0 
	memory_size := []int{} 
	memory_price := []int{} 
	memory_comb := make([]int, len(size))

	for n := 0; n < len(c)-1; n++{ 
		size_value := 0
		price_value := 0
		for i := 0; i < len(c); i++{
			size_value += c[i]*size[i]
			price_value += c[i]*price[i]
		}
		memory_size = append(memory_size, size_value)
		memory_price = append(memory_price, price_value)


		if size_value <= limit{
			if n == 0{
				max_size = size_value
				max_price = price_value
				_ = copy(memory_comb, c)
				memory_slice(max_price, max_size, c)
			}else if n > 0{
				if max_price < memory_price[n]{
					max_size = size_value
					max_price = price_value
					_ = copy(memory_comb, c)
				}
			}
		}
		swap_c := swap(c, n)
		c = nil
		c = swap_c
	}
	return max_size, max_price, memory_comb
}

func main(){
	p := 5 //create first solutions
	result_maxsize := 0
	result_maxprice := 0
	result_minprice := 100
	result_minsize := 0
	result_maxcomb := make([]int, len(size))
	result_mincomb := make([]int, len(size))
	f_comb := [][]int{}

	size, price := sort(size, price)

	g_comb := greedy(size, price)
	f_comb = append(f_comb, g_comb)

	size_v, price_v := []int{}, []int{}
	for i := 0; i < p; i++{
		result_s, result_p, comb := mslsearch(f_comb[0], limit)
		size_v = append(size_v, result_s)
		price_v = append(price_v, result_p)
		if i != 0 && result_p != 0{
			if result_maxprice < price_v[i]{
				result_maxsize = result_s
				result_maxprice = result_p
				_ = copy(result_maxcomb, comb)
			}else if result_minprice > price_v[i]{
				result_minsize = result_s
				result_minprice = result_p
				_ = copy(result_mincomb, comb)
			}
		}else if result_p != 0{
			result_maxsize = result_s
			result_minsize = result_s
			result_maxprice = result_p
			result_minprice = result_p
			_ = copy(result_maxcomb, comb)
			_ = copy(result_mincomb, comb)
		}
		g_comb = shuffle(g_comb)
		f_comb = append(f_comb, g_comb)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("初期解の中での最良な解", memory_max_first, "|", "価値", max_first_p, "|", "容量", max_first_s)
	fmt.Println("初期解の中での最悪な解", memory_min_first, "|", "価値", min_first_p, "|", "容量", min_first_s)
	fmt.Println("--------------------------------------------------")
	fmt.Println("最終解の中での最良な解", result_maxcomb, "|", "価値", result_maxprice, "|", "容量", result_maxsize)
	fmt.Println("最終解の中での最悪な解", result_mincomb, "|", "価値", result_minprice, "|", "容量", result_minsize)
	fmt.Println("--------------------------------------------------")
}