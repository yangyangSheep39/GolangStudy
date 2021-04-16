/*
 * @Description: 
 * @Author: yangyangSheep
 * @Date: 2021-04-16 09:52:22
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-04-16 10:30:55
 */
package main

import "fmt"
/*
 * @Description: 变量使用的注意事项二
 * @Author: yangyangSheep
 * @Date: 2021-04-16 09:52:22
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-04-16 09:58:12
 */
func main() {
	 //该区域的数值可以在同一类型范围内不断变化
	 var i int = 10
	 i = 30
	 i = 50 
	 fmt.Println("i = ",i)
	 //i = "str" //但是不能修改数据类型

	 //在一个函数或在一个代码快中不能重名
	//  var i string = "str"
	//  fmt.Println("i = ",str)

	//程序中+号的使用
	var j int = 20
	var r = i + j	
	fmt.Println("r = " ,r ) // 做加法运算


	 var str string = "str"
	 var str2 string = "str2"
	 var str3 = str + str2	
	 fmt.Println("str3 = " ,str3 ) // 做拼接操作
}


