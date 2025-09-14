package main

import "fmt"


func main() {
	// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
	// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素
   arr := [7]int{1,2,3,4,5,3,2}
   maps := make(map[int]int)
   for _, v := range arr {
      if _, ok := maps[v]; ok {
         maps[v]++
      } else {
         maps[v] = 1
      }
   }

   var oneArr []int = []int{}

   for k, v := range maps {
      if v == 1 {
		oneArr = append(oneArr, k)
      }
   }
   fmt.Println(oneArr)
}