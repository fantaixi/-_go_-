package main

import (
	"fmt"
)

//定义一个结点
type HeroNode struct {
	no        int
	name      string
	nickename string
	pre       *HeroNode //指向上一个结点
	next      *HeroNode //指向下一个结点
}

//给双向链表插入一个结点
//方法1：在双链表的最后插入
func InsertHeroNode(head, newHeroNode *HeroNode) {
	//1、先找到该链表的最后一个结点
	//2、创建一个辅助结点
	temp := head
	for {
		//表示找到最后一个结点
		if temp.next == nil {
			break
		}
		//让temp不断的指向下一个结点
		temp = temp.next
	}
	//3、将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//方法2：根据no的编号从小到大插入
func InsertHeroNode2(head, newHeroNode *HeroNode) {
	//1、找到适当的结点
	//2、创建一个辅助结点
	temp := head
	flag := true
	//让插入的结点的no和temp的下一个结点的no比较
	for {
		//说明到链表的最后
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no { //说明newHeroNode就应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no { //说明链表中已经有这个no，不让插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经有相同的no=", newHeroNode.no)
		return
	} else {
		//先让新的结点指向后面一个结点
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		//再让前面那个结点指向插入的结点
		temp.next = newHeroNode
	}

}

func DelHeroNode(head *HeroNode,id int) {
	temp := head
	flag := false
	for {
		//找到要删除的结点
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			//说明找到
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		//找到，删除
		//直接让结点指向下下一个结点
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	}else {
		fmt.Println("要删除的ID不存在")
	}
}

//显示链表的信息
func ListHeroNode(head *HeroNode) {
	//1、创建一个辅助结点
	temp := head
	//先判断是否为空
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	//2、遍历这个链表
	for {
		//temp.next就表示下一个结点
		fmt.Printf("[%d,%s,%s]==>", temp.next.no, temp.next.name, temp.next.nickename)
		//判断是否链表最后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

//显示链表的信息,逆序打印
func ListHeroNode2(head *HeroNode) {
	//1、创建一个辅助结点
	temp := head
	//先判断是否为空
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	//让temp定位到双向链表的最后
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//2、遍历这个链表
	for {
		//temp.next就表示下一个结点
		fmt.Printf("[%d,%s,%s]==>", temp.no, temp.name, temp.nickename)
		//判断是否到表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}
func main() {
	//1、先创建一个头结点，不需要赋值
	head := &HeroNode{}
	//2、创建一个新的HeroNode
	hero1 := &HeroNode{
		no:        1,
		name:      "宋江",
		nickename: "及时雨",
	}
	hero2 := &HeroNode{
		no:        2,
		name:      "宋江1111",
		nickename: "及时雨",
	}
	hero3 := &HeroNode{
		no:        3,
		name:      "宋江1111",
		nickename: "及时雨",
	}
	hero4 := &HeroNode{
		no:        4,
		name:      "宋江1111",
		nickename: "及时雨",
	}
	//3、加入
	//方法1：
	//InsertHeroNode(head,hero1)
	//InsertHeroNode(head,hero2)
	//方法2：
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero4)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	//4、显示
	ListHeroNode(head)
	fmt.Println()
	DelHeroNode(head,4)
	ListHeroNode(head)
}
