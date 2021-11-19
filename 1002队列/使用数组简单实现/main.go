package main

import (
	"errors"
	"fmt"
	"os"
)

/*
队列是一个有序列表，可以用数组或是链表实现
遵循先入先出
 */

type Queue struct {
	maxSize int
	array [5]int  //数组 => 模拟队列
	front int //表示指向队列首部，但是不含第一个元素
	rear int //表示指向队列的尾部
}

//添加数据
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否满了
	//rear是队列尾部，包含最后一个元素
	if this.rear == this.maxSize-1 {
		return errors.New("queue is full")
	}
	//rear后移
	this.rear++
	this.array[this.rear] = val
	return
}

//从队列取出数据
func (this *Queue) GetQueue() (val int,err error) {
	//先判断队列是否为空
	if this.rear == this.front {
		return -1,errors.New("queue is empty")
	}
	this.front++
	val = this.array[this.front]
	return val,nil
}

//显示队列,找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := this.front+1; i <= this.rear; i++ {
		fmt.Printf("arr[%d]=[%d]\t",i,this.array[i])
	}
	fmt.Println()
}
func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front: -1,
		rear: -1,
	}
	var key string
	var val int
	for  {
		fmt.Println("1.输入add添加")
		fmt.Println("2.输入get取数据")
		fmt.Println("3.输入show显示")
		fmt.Println("4.输入exit退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要入队的数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println("入队成功")
			}
		case "get":
			val,err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println("从队列取出一个数=",val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
