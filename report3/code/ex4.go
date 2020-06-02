package main

import (
	"fmt"
)

var size = []int{3,6,5,4,8,5,3,4,3,5,6,4,8,7,11,8,14,6,12,4}
var price = []int{7,12,9,7,13,8,4,5,3,10,7,5,6,14,5,9,6,12,5,9}
var limit int = 55
var max_size int = 0
var max_price int = 0
var diff_size int = 0
var diff_price int = 0

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

func diff(d_s int, d_p int, n int,s []int, p []int)(int, int){
	diff_num :=  d_s - limit
	diff_s, diff_p := 0, 0
	for l := 0; l < n - 1; l++{
		if s[l] == diff_num {
			diff_s = d_s - s[l]
			diff_p = d_p - p[l]
			break
		}
	}
	return diff_s, diff_p
}

func greedy(s []int, p []int)(int, int){
	for n := 0; n < len(s) - 1; n++{
		max_size += s[n]
		max_price += p[n]
		if max_size > limit{
			diff_s, diff_p := diff(max_size, max_price, n, s, p)
			if diff_size < diff_s{
				diff_size = diff_s
				diff_price = diff_p
			}
			max_size -= s[n]
			max_price -= p[n]
		}
	}
	if diff_price > max_price{
		max_size = diff_size
		max_price = diff_price
	}
	return max_size, max_price
}

func main(){

	size, price := sort(size, price)

	result_size, result_price := greedy(size, price)
	fmt.Println("貪欲法の最大サイズ：", result_size)
	fmt.Println("貪欲法の最大価格：", result_price)
	fmt.Println("貪欲法の最大サイズ_工夫ver：", diff_size)
	fmt.Println("貪欲法の最大価格_工夫ver：", diff_price)
}