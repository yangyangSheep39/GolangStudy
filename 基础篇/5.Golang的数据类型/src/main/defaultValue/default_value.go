package main

import (
	"fmt"
)

/*
 * @Description: 基本数据类型默认值演示
 * @Author: yangyangSheep
 * @Date: 2021-06-10 16:53:34
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-07-27 14:47:37
 */
 func main() {
	var a int //0
	var b float32 //0
	var c float64 //0
	var isMarry bool //0
	var name string //0
	//%f
	fmt.Printf("a=%d,b=%f,c=%f,isMarry=%v,name=%v\n",a,b,c,isMarry,name)
	//%v按照变量的原值输出
	fmt.Printf("a=%d,b=%v,c=%v,isMarry=%v,name=%v",a,b,c,isMarry,name)
 }