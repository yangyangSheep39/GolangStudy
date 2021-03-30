package main

import "fmt" //fmt包中主要提供格式化,输出,输入的函数
/*
 * @Description: Golang的转义字符的使用
 * @Autor: yangyangSheep
 * @Date: 2021-03-20 19:51:45
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-03-30 16:26:10
 */
func main() {
	fmt.Println("制表符: test\tescapeChar")
	fmt.Println("回车: test\nescapeChar")
	fmt.Println("特殊字符\\: test\\escapeChar")
	fmt.Println("特殊字符\": test\"escapeChar")
	fmt.Println("换行符: test\rescapeChar")
}