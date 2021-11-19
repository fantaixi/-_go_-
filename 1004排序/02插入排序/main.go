package main

import "fmt"

/*
左端的数字已完成排序，然后取出那些尚未操作的左端的数字将其与已经操作的左侧的数字进行比较，
如果左边的数字大，交换两个数字，重复此操作，直到出现一个较小的数字或者数字到达左端
*/

func InsertSort(list []int) {
	for i := 1; i < len(list); i++ {
		deal := list[i]
		j := i - 1
		if deal < list[j] {
			for ;j >=0 && deal < list[j];j-- {
				list[j+1] = list[j]
			}
			list[j+1] = deal
		}
	}
}
func main() {
	list := []int{8, 9, 4, 8, 5, 31, 441, 223, 44}
	InsertSort(list)
	fmt.Println(list)
}
