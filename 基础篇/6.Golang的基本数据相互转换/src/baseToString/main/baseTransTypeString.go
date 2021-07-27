package main

import (
	"fmt"
	"strconv"
)

/*
 * @Description: 基本数据类型 转换 string类型
 * @Author: yangyangSheep
 * @Date: 2021-07-27 16:33:58
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-07-27 17:19:57
 */
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string
	
	//使用第一种方式 `fmt.Sprintf("%参数",表达式)`
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q \n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%q \n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q \n",str,str)

	str = fmt.Sprintf("%c",myChar)
	fmt.Printf("str type %T str=%q \n",str,str)

	fmt.Println("================================================")

	//方式二  使用 `strconv` 包的函数  
	str = strconv.FormatInt(int64(num1),10) //第二个参数表示的是一个进制
	fmt.Printf("str type %T str=%q \n",str,str)

	// 'f' 数字按照常规格式输出 'e' 按照科学计数法输出
	//10 表示小数位保留10位 64 表示这个小鼠是float64
	str = strconv.FormatFloat(num2,'f',10,64) 
	fmt.Printf("str type %T str=%q \n",str,str)
	str = strconv.FormatFloat(num2,'e',10,64) 
	fmt.Printf("str type %T str=%q \n",str,str)
	
	str = strconv.FormatBool(b) 
	fmt.Printf("str type %T str=%q \n",str,str)

	fmt.Println("================================================")
	
	//strconv包中有一个很好用的函数Itoa
	var num int = 4576
	str = strconv.Itoa(num) 
	fmt.Printf("str type %T str=%q \n",str,str)
}