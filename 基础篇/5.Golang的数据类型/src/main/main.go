package main

import "fmt"

/*
 * @Description: 数据类型示例
 * @Author: yangyangSheep
 * @Date: 2021-04-16 11:08:44
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-04-16 16:21:47
 */
func main() {
	//整数类型的使用
	var i int = 1
	fmt.Println("i = ", i)

	//测试一下int8的范围
	var minInt8 int8 = -128
	fmt.Println("minInt8 = ",minInt8)
	//会溢出范围
	// var maxInt8Over int8 = 128
	// fmt.Println("maxInt8Over = ",maxInt8Over)

	
	//测试一下uint8(0~255)的范围
	var minUint8 uint8 = 0
	fmt.Println("minUint8 = ",minUint8)
	//会溢出范围
	// var minUint8Over uint8 = -1
	// fmt.Println("minUint8Over = ",minUint8Over)

	//int的其他类型的处理 int uint rune byte(测试时本地系统是64位)
	var a int = 8900
	fmt.Println("a = ",a)

	var b uint = 1 //无符号 报错
	fmt.Println("b = ",b)

	var c byte = 255 //0~255 报错
	fmt.Println("c = ",c)

	var d rune = 255 //int32
	fmt.Println("d = ",d)
}