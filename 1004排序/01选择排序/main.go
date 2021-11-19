package main

import "fmt"

/*
线性搜索数列并找到最小值
 */
func SelectSort(arr *[5]int) {
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i :=j+1; i < len(arr); i++ {
			if max > arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j],arr[maxIndex] = arr[maxIndex],arr[j]
		}
		fmt.Printf("第%d次交换后: %v\n",j+1,*arr)
	}

}
func main() {
	arr := [5]int{10,8,14,5,1}
	SelectSort(&arr)
}
