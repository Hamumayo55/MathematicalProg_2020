package main

import(
	"fmt"
	"math"
	"time"
)

var size = [20]int{3,6,5,4,8,5,3,4,3,5,6,4,8,7,11,8,14,6,12,4}
var price = [20]int{7,12,9,7,13,8,4,5,3,10,7,5,6,14,5,9,6,12,5,9}
var num float64 = 20
var limit int = 55
var max_size int = -1
var max_price int = -1
var result_size = make([]int, 4)
var result_price = make([]int, 4)

func knapsack(start int, end int)(int, int){

	for i := start; i < end-1; i++{
		sum_size := 0
		sum_price := 0
		flag := false
		cal := i

		for j := 0; j < int(num); j++{
			binary := cal%2
			cal = int(math.Floor(float64(cal/2)))
			if binary == 1{
				sum_size += size[j]
				sum_price += price[j]
			}
			if sum_size > limit{
				flag = true
				break
			}
			if flag == false && sum_price > max_price{
				max_size = sum_size
				max_price = sum_price
			}
		}
	}
	return  max_size, max_price
}

func separate(n int)(int, int){
	result_size[0], result_price[0] = knapsack(0,(n/2)-1)
	result_size[1], result_price[1] = knapsack(n/2,n)
	if result_size[0] > result_size[1]{
		return result_size[0], result_price[0]
	}else{
		return result_size[1], result_price[1]
	}
}


func main(){
	n := int(math.Pow(2,num))

	start := time.Now()
	r_size, r_price := separate(n)
	end := time.Now()

	fmt.Printf("処理時間：%f\n", (end.Sub(start)).Seconds())

	fmt.Println("総当たり数：",n-1)
	fmt.Println("最大のサイズ：",r_size)
	fmt.Println("最大の価格：",r_price)
}