package main

import "fmt"
/*
 * @Description: 变量演示案例
 * @Author: yangyangSheep
 * @Date: 2021-04-09 17:00:14
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-04-11 13:14:48
 */

 //声明全局变量
var nWhole1 = 100
var nWhole2 = 200
var nWholeName = "Jack"
//上面的声明方式,也可以改成一次性声明
var	(
	nWholeOneTime1 = 300 
	nWholeOneTime2 = 400
	nWholeOneTimeName = "Nancy"
)
 
func main() {
	//Golang的变量使用方式1
	//第一种:指定变量类型,声明后若不赋值,使用默认值
	//定义变量/声明变量  int 的默认值是0
	var i int;
	fmt.Println("默认值i =" ,i)
	//给i赋值
	i = 10
	//使用变量
	fmt.Println("赋值后i =" ,i)

	//第二种:根据值自行判断变量类型(类型推导)  
	var num = 10.11
	fmt.Println("类型推到后的num =" ,num)

	//第三种:省略`var`,注意 `:=` 左侧的变量不应该是已经声明过的,否则会导致编译报错  
	//这种方式等价于 var name string   name = "Tom"
	// := 的 : 不能省略,否则会报错
	name := "Tom"
	fmt.Println("省略var后的name =" ,name)
	
	//多变量声明:在编程中,有时我们需要一次性声明多个变量,Golang也提供这样的语法  
	var n1,n2,n3 int
	fmt.Println("多变量声明n1 =" ,n1,"多变量声明n2 =" ,n2,"多变量声明n3 =" ,n3)
	//一次性声明多个类型变量
	var n4,nickName,n5 = 100,"Tony",999
	fmt.Println("多类型变量声明n4 =" ,n4,"多类型变量声明nickName =" ,nickName,"多类型变量声明n5 =" ,n5)
	//同样可以使用类型推导
	n6,nameStr,n7 := 100,"Tony",999
	fmt.Println("类型推导变量声明n6 =" ,n6,"类型推导变量声明nameStr =" ,nameStr,"类型推导变量声明n7 =" ,n7)

	//全局变量的打印测试
	fmt.Println("全局变量声明nWhole1 =" ,nWhole1,"全局变量声明nWhole2 =" ,nWhole2,"全局变量声明nWholeName =" ,nWholeName)
	fmt.Println("类型推导变量声明nWholeOneTime1 =" ,nWholeOneTime1,"类型推导变量声明nWholeOneTime2 =" ,nWholeOneTime2,"类型推导变量声明nWholeOneTimeName =" ,nWholeOneTimeName)

}