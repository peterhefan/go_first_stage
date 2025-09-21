package main

import (
	"fmt"
	"strconv"
)

func main() {
	array := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber_1(array))
	fmt.Println(singleNumber_2(array))
	fmt.Println("-------------------------")
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome2(12321))
	fmt.Println("-------------------------")
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([{}])"))
}

func isValid(s string) bool {
	// 括号匹配
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

/**
 * 回文判断一个整数是否是回文数
 */
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		var a = revertedNumber * 10
		b := x % 10
		revertedNumber = a + b
		x = x / 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环结束时我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	if x != revertedNumber && x != revertedNumber/10 {
		return false
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	runes := []rune(strconv.Itoa(x))
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}

	return true
}

func singleNumber_1(array []int) []int {
	// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
	// 找出那个只出现了一次的元素。
	// maps := map[int]int{}
	var maps map[int]int = map[int]int{}
	// maps := make(map[int]int)
	for _, v := range array {
		maps[v]++
	}
	var oneArr []int = []int{}
	for k, v := range maps {
		if v == 1 {
			oneArr = append(oneArr, k)
		}
	}
	return oneArr
}

func singleNumber_2(array []int) int {
	result := 0
	for _, num := range array {
		result ^= num
	}
	return result
}
