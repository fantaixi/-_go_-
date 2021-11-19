package main

import "fmt"

//创建一个结构体
type Boy struct {
	ID int
	Next *Boy
}

//形成一个单向环形链表
//num：要创建多大的链表
//*Boy：返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	//空节点
	first := &Boy{}
	curBoy := &Boy{} //辅助指针
	//判断
	if num < 1 {
		fmt.Println("输入错误")
		return first
	}
	//循环构建环形链表
	for i := 1; i <=num ; i++ {
		boy := &Boy{
			ID: i,
		}
		//如果是第一个
		if i == 1 {
			first = boy  //头结点不能动
			curBoy = boy
			curBoy.Next = first //形成循环
		}else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first  //构成环形
		}
	}
	return first
}

//遍历
func ShowBoy(first *Boy) (n int){
	if first.Next == nil {
		fmt.Println("链表为空。。。")
		return
	}
	//辅助
	curBoy := first
	//求有多少小孩
	num := 1
	for {
		fmt.Printf("小孩编号=%d -->",curBoy.ID)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
		num++
	}
	return num
}

//具体的算法
func PlayGame(first *Boy,startNo int,countNum int)  {
	//处理空
	if first.Next == nil {
		fmt.Println("链表为空。。。。。")
		return
	}
	num := ShowBoy(first)
	if startNo>num {
		fmt.Println("开始数数有错")
		return
	}
	//辅助
	//让tail指向链表的最后一个，在删除的时候需要,此时tail还没有指向最后一个
	tail := first
	for{
		//此时为最后一个
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	//让first移动到startNo[后面删除就以first为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail =tail.Next
	}
	fmt.Println()
	//开始数countNum下，然后就删除first指向的小孩
	for{
		//开始数countNum-1下
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail =tail.Next
		}
		fmt.Printf("小孩编号为%d的出圈 \n",first.ID)
		//删除first指向的小孩
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d的出圈 \n",first.ID)
}
func main() {
	first := AddBoy(5)
	//num := ShowBoy(first)
	//fmt.Println()
	//fmt.Println(num)
	fmt.Println()
	PlayGame(first,2,3)
}
