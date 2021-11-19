package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int //指向队首
	tail    int //指向队尾
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	//tail 在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val //把值给尾部
	this.tail = (this.tail+1)%this.maxSize
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("queue is empty")
	}
	//head指向队列首部，包含第一个元素
	val = this.array[this.head]
	this.head = (this.head+1)%this.maxSize
	return val, nil
}

//显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下：")
	//看队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一个辅助变量指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=[%d]\t", tempHead, this.array[tempHead])
		//让tempHead能回来指向前面的元素
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//判断环形队列满了没有
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判断是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	//*****
	return (this.tail + this.maxSize - this.head) % this.maxSize
}
func main() {
	//先创建一个队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add添加")
		fmt.Println("2.输入get取数据")
		fmt.Println("3.输入show显示")
		fmt.Println("4.输入exit退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要入队的数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("入队成功")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列取出一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
