package main

import "fmt" //fmt包中主要提供格式化,输出,输入的函数
/*
 * @Description: Golang的转义字符的使用
 * @Autor: yangyangSheep
 * @Date: 2021-03-20 19:51:45
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-03-20 19:55:52
 */
func main() {
	fmt.Println("test\tescapeChar")
	fmt.Println("test\nescapeChar")
	fmt.Println("test\\escapeChar")
	fmt.Println("test\"escapeChar")
	fmt.Println("test\rescapeChar")
}