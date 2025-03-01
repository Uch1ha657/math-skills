package main

import (
	"morida/function"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	chi7aja, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("read error file")
		return
	}

	arr := string(chi7aja)
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

	avg := function.Average(nums)
	md := function.Median(nums)
	vari := function.Variance(nums)
	Stand := function.Standard(nums)
	fmt.Println("avg:", int(math.Round(avg)))
	fmt.Println("md:", int(math.Round(md)))
	fmt.Println("vari:", int(math.Round(vari)))
	fmt.Println("Stand", int(math.Round(Stand)))
}
