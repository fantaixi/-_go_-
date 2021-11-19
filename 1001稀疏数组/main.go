package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	Rows int
	Cols int
	Value int
}

func main() {
	//1、创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	chessMap[2][2] = 1
	chessMap[3][2] = 2
	chessMap[4][2] = 1
	chessMap[4][1] = 2

	//2、输出数组
	fmt.Println("原始数组为：")
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------------")
	//3、转成稀疏数组
	//1）遍历chessMap，如果有一个元素的值不为0，创建一个结构体
	//2）将其放入对应的切片
	var sparseArr []ValNode
	//一个标准的稀疏数组应该还有原始数组的行和列，默认值
	valNode := ValNode{
		Rows: 11,
		Cols: 11,
		Value: 0,
	}
	sparseArr = append(sparseArr,valNode)
	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			//如果有值不为0
			if v2!=0 {
				//创建一个ValNode值节点
				valNode := ValNode{
					Rows: i,
					Cols: j,
					Value: v2,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("稀疏数组为：")
	for _, valNode := range sparseArr {
		fmt.Printf("%d %d %d\n",valNode.Rows,valNode.Cols,valNode.Value)
	}

	//写入文件
	filePath := "./spare_arr.data"
	file,err := os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,0666)
	if err != nil {
		fmt.Println("Oper file err",err.Error())
		return
	}
	defer file.Close()
	writer:= bufio.NewWriter(file)
	for _, v := range sparseArr {
		writer.WriteString(fmt.Sprintf("%d %d %d \n",v.Rows,v.Cols,v.Value))
	}
	writer.Flush()

	//从文件中恢复
	ReadSparseArray(filePath)
}

func ReadSparseArray(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()
	bfrd := bufio.NewReader(file)
	var index = 0
	var arr [][]int
	for {
		line, err := bfrd.ReadBytes('\n')
		if err != nil {
			break
		}
		index++
		temp := strings.Split(string(line), " ")
		row, _ := strconv.Atoi(temp[0])
		col, _ := strconv.Atoi(temp[1])
		value, _ := strconv.Atoi(temp[2])
		if index == 1 {
			for i := 0; i < row; i++ {
				var arr_temp []int
				for j := 0; j < col; j++ {
					arr_temp = append(arr_temp, value)
				}
				arr = append(arr, arr_temp)
			}
		}
		if index != 1 {
			arr[row][col] = value
		}
	}
	//fmt.Println(arr)
	// 打印数据
	fmt.Println("从磁盘读取后的数据")
	for _, v := range arr {
		for _, v1 := range v {
			fmt.Printf("%d\t", v1)
		}
		fmt.Println()
	}
}
