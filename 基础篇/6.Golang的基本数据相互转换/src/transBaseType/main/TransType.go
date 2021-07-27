package main

import (
	"fmt"
)

/*
 * @Description: 用于演示Golang的基本数据类型的转换以及细节演示
 * @Author: yangyangSheep
 * @Date: 2021-07-27 14:59:01
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-07-27 15:56:06
 */
func main() {
	/*Golang的基本数据类型的转换*/
	var i int = 100
	//希望将i转成float类型
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	// var n2 float32 = i
	fmt.Printf("i=%v n1=%v n2=%v \n",i,n1,n2)

	/*细节演示*/
	//Golang中,数据类型的转换是可以从 **表示范围小 -> 表示范围大**,也可以 **表示范围大 -> 表示范围小**  
	//被转换的是**变量存储的数据(即值)**,变量本身的数据类型并没有变化  
	//在转换中,比如将int64转换成int8[-128~+127],编译时不会报错,只是转换的结果是按溢出处理,和我们希望的结果不一样,因此在转换时需要考虑范围的问题
	fmt.Printf("i type is %T \n",i)
	var num1 int64 = 999999
	var num2 int8 = int8(num1) //编译不会报错,溢出处理(按照二进制的处理)
	fmt.Printf("num1=%v num2=%v \n",num1,num2)
}