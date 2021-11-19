package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
创建一个数组模拟队列，每隔一定时间（随机），给该数组添加一个数
启动两个协程，每隔一定时间（随机）到队列取出数据
在控制台输出：
X号协程 服务 --> X号客户
。。。
使用锁机制
 */
type Queue struct {
	maxSize int
	array [15]int
	head    int //指向队首
	tail    int //指向队尾
}
// 声明一个全局互斥锁
var lock sync.Mutex

//给队列添加数据
func (this *Queue) Add(num int) (err error) {
	//先判断队列是否已满
	if this.tail == this.maxSize-1 {
		return errors.New("Queue is Full")
	}
	this.tail++
	this.array[this.tail] = num
	return
}

//每隔一段随机时间添加数据到队列
func (this *Queue) AddRandomTime() {
	num := 0
	randNum := 0
	//设置随机源
	rand.Seed(time.Now().UnixNano())
	for {
		//让程序随机睡眠一段时间
		randNum = rand.Intn(500)
		time.Sleep(time.Millisecond * time.Duration(randNum))
		num++
		//往队列添加数据
		this.Add(num)
		if this.tail == this.maxSize-1 {
			break
		}
	}
}
//出队列
func (this *Queue) Get() (value int, err error) {
	//先判断队列是否为空
	if this.tail == this.head {
		return -1,errors.New("Queue is Empty")
	}
	this.head++
	value = this.array[this.head]
	return value,err
}

//取出环形队列有多少个元素
func (this *Queue) Size() int {
	return (this.tail+this.maxSize-this.head)%this.maxSize
}
//每隔一段随机时间从队列取数据
func (this *Queue) GetRandomTime(id int)  {
	//设置随机源
	rand.Seed(time.Now().UnixNano())
	for {
		//让程序随机睡眠一段时间
		randNum := rand.Intn(500)
		time.Sleep(time.Millisecond * time.Duration(randNum))
		//上锁
		lock.Lock()
		num,err := this.Get()
		lock.Unlock()
		if err != nil {
			fmt.Println("队列为空")
		}else {
			fmt.Printf("%d 号协程服务--->%d号客户\n",id,num)
		}
	}
	
}
func main() {
	queue := &Queue{
		maxSize: 15,
		head: -1,
		tail: -1,
	}
	go queue.AddRandomTime()
	for i := 1; i <= 2; i++ {
		go queue.GetRandomTime(i)
	}
	time.Sleep(time.Second * 1000)
}
