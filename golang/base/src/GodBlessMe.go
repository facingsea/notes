package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	runShuangSeQiu()
}

/**
 * 模拟双色球
 * @return {nil}
 */
func runShuangSeQiu() {

	for i := 0; i < 10; i++ { // 生成10注
		// red:33, blue: 16
		// red
		var redBalls [6]int
		for j := 0; j < 6; j++ {
			red := rand.Intn(33)
			is := has(redBalls[0:], red)
			if is {
				j--
			} else {
				redBalls[j] = red
			}
		}

		sortBalls := sort(redBalls[0:])

		for _, v := range sortBalls {
			//fmt.Print(v)
			if v >= 10 {
				fmt.Print("  " + strconv.Itoa(v))
			} else {
				fmt.Print("  0" + strconv.Itoa(v))
			}
		}

		fmt.Print(" , ")

		var blueBalls [1]int
		for j := 0; j < 1; j++ {
			blue := rand.Intn(16)
			is := has(blueBalls[0:], blue)
			if is {
				j--
			} else {
				blueBalls[j] = blue
			}
		}

		for _, v := range blueBalls {
			//fmt.Print(v)
			if v >= 10 {
				fmt.Print("  " + strconv.Itoa(v))
			} else {
				fmt.Print("  0" + strconv.Itoa(v))
			}
		}

		// blue

		fmt.Println("")
	}

}

func has(arr []int, val int) (is bool) {
	if cap(arr) == 0 {
		return false
	}
	for i := 0; i < cap(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

func sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
