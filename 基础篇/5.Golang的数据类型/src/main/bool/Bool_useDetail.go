package main

import (
	"fmt"
	"unsafe"
)

/*
 * @Description: bool类型使用演示
 * @Author: yangyangSheep
 * @Date: 2021-06-02 11:04:40
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-06-02 17:42:53
 */
 func main() {
	var b = false
	fmt.Println("b=",b)
	
	//注意事项 
	//1.bool类型占用的空间是一个字节
	fmt.Println("b的占用空间 = ",unsafe.Sizeof(b))
	//2.bool类型只能取true或者false
	// b = -1
 }	
