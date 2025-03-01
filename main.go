package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	math, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("read error file")
		return
	}

	arr := string(math)
	nums := []int{}
	words := strings.Fields(arr)
	for _, v := range words {
		arrr := v
		num, err := strconv.Atoi(arrr)
		if err != nil {
			fmt.Println("eroore help", err)
		}

		nums = append(nums, num)
		
	}
	
	avg := Average(nums)
	md := Median(nums)
	vari := Variance(nums)
	Stand := Standard(nums)
	
	fmt.Println("avg:", avg)
	fmt.Println("md:", md)
	fmt.Println("vari:", vari)
	fmt.Println("Stand", Stand)

}

func Average(nums []int) float64 {
	lmjmou3 := 0

	for _, v := range nums {
		lmjmou3 += v

	}
	llen := len(nums)
	res := lmjmou3 / llen
	return float64(res)
}

func Median(nums []int) float64 {
	sort.Ints(nums)
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	} else {
		n1 := nums[len(nums)/2-1]
		n2 := nums[(len(nums) / 2)]
		return float64((n1 + n2) / 2)
	}
}
 func Variance(nums []int) float64{
	avg := Average(nums)
	rest:= 0
	result := 0
	for _, v := range nums {
		res:= v - int(avg) 
		rest += res * res 
		 
	 result = rest/len(nums)
	}
	return float64(result)
 }
 func Standard(nums []int) float64{
	Sta:= Variance(nums)
	result := math.Sqrt(float64(Sta))
	return result

 }