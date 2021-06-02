package main

// import "fmt"
// import "unsafe"
//上述的引入包方式不够简洁
import (
	"fmt"
	"unsafe"
)
/*
 * @Description: 整型的使用细节
 * @Author: yangyangSheep
 * @Date: 2021-04-16 16:21:18
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-06-02 17:38:46
 */
 func main() {
	 //整型的使用细节
	 //查看变量的数据类型和所占用的字节数
	var n1 = 100 //此处n1是什么类型怎么查看
	// fmt.Printf("", var) 可以用于做格式化输出
	fmt.Printf("n1 的类型是 %T \n",n1)
	
	var n2 = 1000 //此处n1是什么类型怎么查看
	// fmt.Printf("", var) 可以用于做格式化输出
	fmt.Printf("n2 的类型是 %T 占用的字节数是 %d \n",n2,unsafe.Sizeof(n2))
	
	//保小不保大的原则
	//在保证程序正确运行下,尽量使用占用空间小的数据类型[如年龄]
	var ageUnbelievable int64 = 150
	var age byte = 150
	fmt.Printf("ageUnbelievable 的类型是 %T 占用的字节数是 %d \n",ageUnbelievable,unsafe.Sizeof(ageUnbelievable))
	fmt.Printf("age 的类型是 %T 占用的字节数是 %d \n",age,unsafe.Sizeof(age))
 }	
