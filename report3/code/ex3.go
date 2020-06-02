package main

import (
	"fmt"
)

var size = []int{3,6,5,4,8,5,3,4}
var price = []int{7,12,9,7,13,8,4,5}
var limit int = 25
var max_size int = 0
var max_price int = 0

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

func greedy(s []int, p []int)(int, int, []int){
	comb := []int{}
	for n := 0; n < len(s) - 1; n++{
		max_size += s[n]
		max_price += p[n]
		if max_size > limit{
			max_size -= s[n]
			max_price -= p[n]
			comb = append(comb, 0)
		}else{
			comb = append(comb, 1)
		}
	}
	return max_size, max_price, comb
}

func main(){

	size, price := sort(size, price)

	result_size, result_price, result_comb := greedy(size, price)
	fmt.Println("組み合わせ：", result_comb)
	fmt.Println("貪欲法の最大サイズ：", result_size)
	fmt.Println("貪欲法の最大価格：", result_price)
}