package main

import(
	"fmt"
	"unsafe"
)
/*
 * @Description: 浮点类型的使用细节
 * @Author: yangyangSheep
 * @Date: 2021-04-16 16:59:21
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-06-02 17:43:26
 */
func main() {
	var price float32 = 89.123
	fmt.Println("price = ",price)

	//浮点数都是有符号位的,说明浮点数都有负数
	var f1 float32 = -1324.4616
	fmt.Println("f1 = ",f1)

	var f2 float64 = -54874651.464165 
	fmt.Println("f2 = ",f2) //科学计数法显示 -5.4874651464165e+07 = 5.4874651464165 * 10^7

	//尾数部分可能丢失,造成精度损失
	var f3 float32 = -123.0000901
	var f4 float64 = -123.0000901
	fmt.Println("f3 = ",f3)
	fmt.Println("f4 = ",f4)



	//Golang的浮点类型默认声明为float64类型
	// Golang的浮点类型固定的范围和字段长度,不受具体OS的影响
	var f5  = 8456.464
	fmt.Printf("f5 的类型是 %T 占用的字节数是 %d \n",f5,unsafe.Sizeof(f5))


	// 十进制数形式,如 5.12 0.512(必须有小数点)  
	// 浮点类型常量有两种表示形式  
	// 科学计数法形式,如 2.1234e2 = 2.12 * 10^2 5.12E-2=5.12/10^2  
	fl6 := 5.12
	fl7 := .512
	fmt.Println("fl6 = ",fl6)
	fmt.Println("fl7 = ",fl7)
	
	
	fl8 := 5.1234e2  // 5.1234 * 10^2
	fmt.Println("fl8 = ",fl8)
	fl9 := 5.1234E2  // 5.1234 * 10^2
	fmt.Println("fl9 = ",fl9)
	fl10 := 5.1234E-2 // 5.1234 / 10^2
	fmt.Println("fl10 = ",fl10)
}