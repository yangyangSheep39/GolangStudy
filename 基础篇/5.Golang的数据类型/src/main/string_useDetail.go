package main

import (
	"fmt"
)

/*
 * @Description: 字符串的使用细节
 * @Author: yangyangSheep
 * @Date: 2021-06-02 14:30:40
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-06-02 16:28:12
 */
 func main() {
	//string的基本使用
	var address string = "安徽芜湖"
	fmt.Println(address)

	//使用细节
	//字符串不可变
	var str = "hello"
	// str[0] = 'a' //这里不能去修改str的内容 Go中的字符串是不能变的
	fmt.Println(str)

	//表现形式
	// 1. 双引号: 会是别转义字符  
	var str2 = "abc\nefg"
	fmt.Println(str2)
	
	// 2. 反引号: 以字符串的原生形式输出,包括换行和特殊字符,可以实现防止攻击/输出源代码等效果  
	str3 := `//string的基本使用
	var address string = "安徽芜湖"
	fmt.Println(address)

	//使用细节
	//字符串不可变
	var str = "hello"
	// str[0] = 'a' //这里不能去修改str的内容
	fmt.Println(str)

	//表现形式
	// 1. 双引号: 会是别转义字符  
	var str2 = "abc\nefg"
	fmt.Println(str2)
	
	// 2. 反引号: 以字符串的原生形式输出,包括换行和特殊字符,可以实现防止攻击/输出源代码等效果`
	fmt.Println(str3)
	
	//字符串拼接
	str4 := "hello" + "world"
	str4 += "nihao"
	fmt.Println("str4= ",str4)

	//当一个拼接的操作很长时,可以分行写,但是要把加号留在上面
	str5 := "hello" + "world" +
	"hello" + "world" +
	"hello" + "world" +
	"hello" + "world" + "hello" + "world"
	fmt.Println("str5= ",str5)
}	
