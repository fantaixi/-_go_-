package main

import "fmt"

/*
第一个操作的对象是数列中所有的数字，选择一个数字作为排序的基准，这个数字称为pivot，
pivot通常随机一个数字（一般选择最左或最右），在pivot上做一个标记，接下来，
在最左边的数字上标记左标记，最右边的数字标记右标记，将左边的标记向右移动，当左标记到达超过pivot的数字是，停止移动，
继续讲右标记向右移动，当右标记到达小于pivot的数字时，停止移动，当左右的标记停止时，交换标记的数字，之后不停的递归
*/

func QuickSort(left int, right int, list []int) {
	l := left  //左边的下标
	r := right  //右边的下标
	//中轴
	pivot := list[(l+r)/2]
	temp := 0
	for ;l<r; {
		for ; list[l] < pivot; {
			l++
		}
		for ; list[r] > pivot; {
			r--
		}
		if l >= r {
			break
		}
		temp = list[l]
		list[l] = list[r]
		list[r] = temp
		if list[l] == pivot {
			r--
		}
		if list[r] == pivot {
			l++
		}
	}
	//如果l=r，就再移动一下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left,r,list)
	}
	//向右递归
	if right > l {
		QuickSort(l,right,list)
	}
}
func main() {
	list := []int{5,9,7,6,3,1,4}
	QuickSort(0,len(list)-1,list)
	fmt.Println(list)
}
