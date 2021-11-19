package main

import "fmt"

//myMap *[8][7]int:地图
// i , j 表示对哪个点进行测试
func SetWay(myMap *[8][7]int, i, j int) bool{
	//如果找到myMap[6][5],就是成功
	if myMap[6][5] == 2 {
		return true
	}else {
		if myMap[i][j] == 0 {  //表示这个点可以探测
			myMap[i][j] = 2
			//开始探路
			//使用下右上左策略
			if  SetWay(myMap, i+1, j) { //向下
				return true
			}else if SetWay(myMap, i, j+1) { //向右
				return true
			}else if  SetWay(myMap, i-1, j) { //向上
				return true
			}else if SetWay(myMap, i, j-1) { //向左
				return true
			}else { //说明是死路
				myMap[i][j] = 3
				return false
			}
		}else { //说明这个点是1，是墙
			return false
		}
	}
}
func main() {
	//用二维数组模拟迷宫
	// 元素的值为1表示墙
	// 元素的值为2表示一个通路
	// 元素的值为0表示没有走过的点
	// 元素的值为3表示走过的点，但是走不通
	var myMap [8][7]int
	//先把地图的上下左右设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[3][3] = 1
	fmt.Println("地图为：")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j]," ")
		}
		fmt.Println()
	}
	fmt.Println("探测完毕之后：")
	SetWay(&myMap,1,1)
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j]," ")
		}
		fmt.Println()
	}
}
