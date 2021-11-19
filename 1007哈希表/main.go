package main

import (
	"fmt"
	"os"
)

//定义emp
type Emp struct {
	ID int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到员工 %d\n",this.ID %7,this.ID)
}
//定义EmpLink
//不带表头，第一个节点就存放数据
type EmpLink struct {
	Head *Emp
}
//添加员工的方法，保证添加时编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil  //pre在cur之前
	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp
		return
	}
	//如果不是空链表，给emp找到对应的位置并插入
	//让cur和emp比较，然后让pre保持在cur前面
	for{
		if cur != nil {
			//比较
			if cur.ID > emp.ID {
				//找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		}else {
			break
		}
	}
	//退出时，判断是否将emp添加到链表的最后
	pre.Next = emp
	emp.Next = cur
}

func (this *EmpLink) ShowLink(id int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n",id)
		return
	}
	//显示数据
	cur := this.Head
	for  {
		if cur != nil {
			fmt.Printf("链表%d 员工ID=%d 名字=%s -->",id,cur.ID,cur.Name)
			cur = cur.Next
		}else {
			break
		}
	}
	fmt.Println()
}

func (this *EmpLink) Find(id int) *Emp {
	cur := this.Head
	for{
		if cur != nil && cur.ID == id {
			return cur
		}else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}
//定义HashTable，含有链表的数组
type HashTable struct {
	LinkArr [7]EmpLink
}
//HashTable 的 add
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将员工添加到哪个链表
	linkNo := this.HashFun(emp.ID)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}
//编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7  //得到对应的链表的下标
}
//显示
func (this *HashTable) Show() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}
//查找
func (this *HashTable) FindByID(id int) *Emp {
	//确定员工应该在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].Find(id)
}
func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for{
		fmt.Println("==========菜单==========")
		fmt.Println("input 添加员工")
		fmt.Println("show 显示员工")
		fmt.Println("find 查找员工")
		fmt.Println("exit 退出")
		fmt.Println("输入选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("员工ID")
			fmt.Scanln(&id)
			fmt.Println("员工名字")
			fmt.Scanln(&name)
			emp := &Emp{
				ID: id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.Show()
		case "find":
			fmt.Println("请输入要查找的id：")
			fmt.Scanln(&id)
			emp := hashTable.FindByID(id)
			if emp == nil {
				fmt.Printf("id=%d 的员工不存在\n",id)
			}else {
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
