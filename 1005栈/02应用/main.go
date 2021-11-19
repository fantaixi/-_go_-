package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int   //可以存放多少个数
	Top    int   //表示栈顶，因为栈顶固定，可以直接使用
	arr    [20]int //数组模拟栈
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
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	//判断栈是否为空
	if this.Top == -1 {
		fmt.Println("栈为空。。。")
		return 0, errors.New("stack full...")
	}
	//先取值，再让top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//判断一个字符是否是运算符（+-*/）
func (this *Stack) IsOper(val int) bool {
	//ascII 码
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (this *Stack) Cal(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符有错。。")
	}
	return res
}

//返回运算符的优先级
//假设 乘和除的优先级是1  加和减的优先级是0
func (this *Stack) Prio(oper int) int {
	if oper == 42 || oper == 47 {
		return 1
	}else {
		return 0
	}
}
func main() {
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "3000+3*6-90000"
	//帮助扫描exp
	index := 0
	num1 :=0
	num2 :=0
	oper := 0
	result := 0

	keepNum := ""
	for {
		ch := exp[index : index+1] //字符串
		//转成对应的数字
		temp := int([]byte(ch)[0])
		//说明是符号
		if operStack.IsOper(temp) {
			//如果是空栈，直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			}else {
				if operStack.Prio(operStack.arr[operStack.Top]) >= operStack.Prio(temp) {
					num1,_ = numStack.Pop()
					num2,_ = numStack.Pop()
					oper,_ = operStack.Pop()
					result = operStack.Cal(num1,num2,oper)
					//再把求出的值压入栈
					numStack.Push(result)
					//当前的符号压入符号栈
					operStack.Push(temp)
				}else {
					operStack.Push(temp)
				}
			}
		}else {  //说明是数
			//处理多位数
			//用keepNum做拼接，每次向index前面字符测试一下，看是否是运算符，然后处理
			keepNum += ch
			if index == len(exp)-1 {
				val,_ := strconv.ParseInt(keepNum,10,64)
				numStack.Push(int(val))
			}else{
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val,_ := strconv.ParseInt(keepNum,10,64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
			//val,_ := strconv.ParseInt(ch,10,64)
			//numStack.Push(int(val))
		}
		//继续扫描
		if index+1 == len(exp) {
			break
		}
		index++
	}
	for{
		if operStack.Top == -1 {
			break
		}
		num1,_ = numStack.Pop()
		num2,_ = numStack.Pop()
		oper,_ = operStack.Pop()
		result = operStack.Cal(num1,num2,oper)
		//再把求出的值压入栈
		numStack.Push(result)
	}
	//如果没有错误，结果就是数栈的最后一个
	res,_ := numStack.Pop()
	fmt.Println(res)
}
