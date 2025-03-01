package function

import (
	"math"
	"sort"
)

func Average(nums []int) float64 {
	lmjmou3 := 0

	for _, v := range nums {
		lmjmou3 += v
	}
	llen := len(nums)
	res := float64(lmjmou3) / float64(llen)
	return res
}
func Median(nums []int) float64 {
	sort.Ints(nums)
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	} else {
		n1 := nums[len(nums)/2-1]
		n2 := nums[(len(nums) / 2)]
		return ((float64(n1) + float64(n2)) / 2)
	}
}
func Variance(nums []int) float64 {
	avg := Average(nums)
	rest := 0.0
	for _, v := range nums {
		res := float64(v) - avg
		rest += res * res

	}
	result := rest / float64(len(nums))
	return float64(result)
}
func Standard(nums []int) float64 {
	Sta := Variance(nums)
	result := math.Sqrt(float64(Sta))
	return result
}
