//一个go文件需要在一个包中
package main

import "fmt"
/*
 * @Description: Go的特性快速体验
 * @Autor: yangyangSheep
 * @Date: 2021-03-19 19:20:29
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-03-20 19:15:12
 */
 //这里需要使用main函数去运行
 func main(){
	 //输出一句话
	 fmt.Println("Hello World")
 }

 //写一个函数，实现同时返回 “和” 和 “差”
//go函数可以同时支持返回多个值
 func getSumAndSub(n1 int,n2 int )(int,int){
	sum = n1 + n2
	sub = n1 - n2
	return sum ,sub
 } 
 
