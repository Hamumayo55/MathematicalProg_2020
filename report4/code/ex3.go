package main

import (
	"fmt"
	"math/rand"
	"time"
)

var size = []int{3,6,5,4,8,5,3,4}
var price = []int{7,12,9,7,13,8,4,5}
// 最小・最大判定
var min_first_p int = 1000 
var max_first_p int = 0
var min_first_s int = 0 
var max_first_s int = 0
var memory_min_first = make([]int, len(size)) //初期解の中の最良解
var memory_max_first = make([]int, len(size)) //初期解の中の最悪解
var first_flag bool = true

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

func create_binary(p int) [][]int{ //ランダムに初期解を作る関数
	rand.Seed(time.Now().Unix())
	binary := [][]int{}
	slice := []int{}
	for n := 0; n < p; n++{
		for i := 0; i < len(size); i++{
			slice = append(slice, rand.Intn(2))
		}
		binary = append(binary, slice)
		slice = nil
	}
	return binary
}

//扇動関数
func swap(c []int, i int) []int{
	swap_c := c
	swap_c[i], swap_c[i+1] =  swap_c[i+1], swap_c[i]
	return swap_c
}

//初期解の中の最良と最悪の解を判定する解
func memory_slice(f_p int, f_s int,c []int){ 
	if first_flag {
		min_first_p, max_first_p = f_p, f_p
		min_first_s, max_first_s = f_s, f_s
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
	limit := 25
	p := 5 //初期解を生成する数
	result_maxsize := 0
	result_maxprice := 0
	result_minprice := 0
	result_minsize := 0
	result_maxcomb := make([]int, len(size))
	result_mincomb := make([]int, len(size))

	comb := create_binary(p)
	size_v, price_v := []int{}, []int{}
	for i := 0; i < p; i++{
		result_s, result_p, comb := mslsearch(comb[i], limit)
		size_v = append(size_v, result_s)
		price_v = append(price_v, result_p)
		if i != 0{
			if result_maxprice < price_v[i]{
				result_maxsize = result_s
				result_maxprice = result_p
				_ = copy(result_maxcomb, comb)
			}else if result_minprice > price_v[i]{
				result_minsize = result_s
				result_minprice = result_p
				_ = copy(result_mincomb, comb)
			}
		}else {
			result_maxsize = result_s
			result_minprice = result_p
			result_maxprice = result_p
			result_minprice = result_p
			_ = copy(result_maxcomb, comb)
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("初期解の中での最良な解", memory_max_first, "|", "価値", max_first_p, "|", "容量", max_first_s)
	fmt.Println("初期解の中での最悪な解", memory_min_first, "|", "価値", min_first_p, "|", "容量", min_first_s)
	fmt.Println("--------------------------------------------------")
	fmt.Println("最終解の中での最良な解", result_maxcomb, "|", "価値", result_maxprice, "|", "容量", result_maxsize)
	fmt.Println("最終解の中での最悪な解", result_mincomb, "|", "価値", result_minprice, "|", "容量", result_minsize)
	fmt.Println("--------------------------------------------------")
}