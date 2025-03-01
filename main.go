package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"morida/function"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run . data.txt")
		return
	}
	data := os.Args[1]
	chi7aja, err := os.ReadFile(data)
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
