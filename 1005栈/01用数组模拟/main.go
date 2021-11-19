package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int //可以存放多少个数
	Top int //表示栈顶，因为栈顶固定，可以直接使用
	arr [5]int //数组模拟栈
}

//入栈
func (this *Stack) Push(val int) (err error) {
	//先判断是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满了。。。")
		return errors.New("stack full...")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//遍历，从栈顶开始遍历
func (this *Stack) List() {
	//先判断是否为空
	if this.Top == -1 {
		fmt.Println("栈为空。。。")
		return
	}
	for i := this.Top; i >=0; i-- {
		fmt.Printf("arr[%d]=%d\n",i,this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	//判断栈是否为空
	if this.Top == -1 {
		fmt.Println("栈为空。。。")
		return 0,errors.New("stack full...")
	}
	//先取值，再让top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}
func main() {
	 stack := &Stack{
		 MaxTop: 5,
		 Top: -1,
	 }
	 //入栈
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	stack.List()
	//出栈
	val,_:= stack.Pop()
	fmt.Println(val)
	stack.List()
}
