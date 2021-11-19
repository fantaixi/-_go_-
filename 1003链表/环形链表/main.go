package main

import "fmt"

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertCatNode(head, newCatNode *CatNode) {
	//判断是否添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name=newCatNode.name
		head.next = head //形成环状，这里是自己指向自己
		fmt.Println(newCatNode,"加入到环形链表")
		return
	}
	//定义一个临时变量，辅助找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表
	temp.next = newCatNode
	newCatNode.next = head
}

//遍历
func ListCatNode(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("此时为空")
		return
	}
	for{
		fmt.Printf("猫的信息为=[no:%d,name:%s]==>\n",temp.no,temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除一个结点
func DelCatNode(head *CatNode,id int) *CatNode{
	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("环形链表为空，无法删除")
		return head
	}
	//如果只有一个结点
	if temp.next == head {
		temp.next = nil
		return head
	}
	//将helper定位到链表最后
	for{
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	//如果有两个或以上的结点
	for{
		//如果到这里，说明比较到最后一个，但是最后一个还没比较，出去再比较
		if temp.next == head {
			break
		}
		//如果找到要删除的id
		if temp.no == id {
			if temp==head { //说明删除的是头结点
				head = head.next
			}
			//删除
			helper.next = temp.next
			fmt.Printf("要删除的猫的id=%d\n",id)
			flag = false
			break
		}
		temp = temp.next   //temp用于比较
		helper = helper.next   //helper用于干掉temp
	}

	//这里还要再比较一次
	//如果flag为真，则在上面没有删除过
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫=%d\n",id)
		}else {
			fmt.Printf("没有这只猫=%d\n",id)
		}
	}
	return head
}
func main() {
	//初始化一个环形链表的头结点
	head := &CatNode{}
	cat1 := &CatNode{
		no: 1,
		name: "tom",
	}
	cat2 := &CatNode{
		no: 2,
		name: "tom",
	}
	cat3 := &CatNode{
		no: 3,
		name: "tom",
	}
	InsertCatNode(head,cat1)
	InsertCatNode(head,cat2)
	InsertCatNode(head,cat3)
	ListCatNode(head)
	head = DelCatNode(head,30)
	ListCatNode(head)


}
