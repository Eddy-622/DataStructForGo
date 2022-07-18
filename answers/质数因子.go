package main

import "fmt"
import "strconv"

func main() {
	var s string
	fmt.Scan(&s)
	num, _ := strconv.Atoi(s)
	if num <= 1 {
		return
	}
	// 从2开始往后找
	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			fmt.Print(i)
			fmt.Print(" ")
			num /= i
		}
	}
	if num >= 2 {
		fmt.Print(num)
	}
}
