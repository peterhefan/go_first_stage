package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// array := []int{4, 1, 2, 3, 1, 2}

	// fmt.Println(singleNumMatch(array))
	// fmt.Println(singleNumber_1(array))
	// fmt.Println(singleNumber_2(array))
	// fmt.Println("-------------------------")
	// fmt.Println(isPalindrome(12321))
	// fmt.Println((12321))
	// fmt.Println("-------------------------")
	// fmt.Println(isValid("()[]{}"))
	// fmt.Println(isValid("([{}])"))

	// fmt.Println(isValid2("()[]{}"))
	// fmt.Println(isValid2("[]{()}"))
	// fmt.Println(isValid2("([{}])"))

	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longCommonPrefix(strs))
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longCommonPrefix(strs2))

	testCases := [][]int{
        {1, 2, 3},
        {4, 3, 2, 1},
        {9},
        {9, 9, 9},
        {0},
    }
    for _, tc := range testCases {
        fmt.Printf("输入: %v\n", tc)
        fmt.Printf("输出: %v\n\n", plusOne(tc))
    }
}


// 基本值类型
// 加一 
// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。
// 这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func plusOne(digits []int) []int {

	n := len(digits)
	for i := n-1 ; i >=0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1},digits...)
}


// 最长公共前缀
func longCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix

}

// 有效的括号
func isValid2(str string) bool {

	maps := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := []rune{}
	for _, char := range str {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if maps[char] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println("stack:", stack, "")

	return len(stack) == 0
}

// 有效的括号
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

/**
 * 回文判断一个整数是否是回文数
 */
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

// 只出现一次的数字
func singleNumMatch(arr []int) []int {

	maps := map[int]int{}
	// maps := make(map[int]int)
	for _, v := range arr {
		// fmt.Println(index)
		if maps[v] != 0 {
			maps[v]++
		} else {
			maps[v] = 1
		}
	}

	result := []int{}
	for k, v := range maps {
		// fmt.Println(a)
		if v == 1 {
			// fmt.Print(k, ",")
			result = append(result, k)
		}
	}

	// 升序
	sort.Ints(result)

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result

}

// 只出现一次的数字
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

// 只出现一次的数字
func singleNumber_2(array []int) int {
	result := 0
	for _, num := range array {
		result ^= num
	}
	return result
}
