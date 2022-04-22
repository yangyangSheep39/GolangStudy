package main

import(
	"fmt"
)

/*
 * @Description: 指针类型演示
 * @Author: yangyangSheep
 * @Date: 2022-04-22 14:36:18
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2022-04-22 15:33:27
 */
func main()  {
	//基本数据类型在内存的布局
	var i int = 10
	//取i的地址
	fmt.Println("i的地址=",&i)

	//1.ptr是一个指针变量
	//2.ptr的类型*int
	//3.ptr本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Println("ptr的地址=",&ptr)
	fmt.Printf("ptr指向的值=%v\n",*ptr)
}