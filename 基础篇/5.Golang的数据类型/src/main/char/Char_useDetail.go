package main


// import "fmt"
// import "unsafe"
//上述的引入包方式不够简洁
import (
	"fmt"
)
/*
 * @Description: 字符类型的使用细节
 * @Author: yangyangSheep
 * @Date: 2021-06-01 15:31:19
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2021-06-02 17:43:04
 */
 func main() {
	var c1 byte = 'a' 
	var c2 byte = '0' //字符的0

	//当我们直接输出byte值的时候,就是输出对应的字符的码值
	//'a' => 97  '0' => 48
	fmt.Println("c1=",c1)
	fmt.Println("c2=",c2)
	//如果我们希望输出对应字符,需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2)
	fmt.Println("=========================")

	// var c3 byte  = '北' //他的字符不在byte的范围中 overflow溢出
	var c3 int  = '北' //他的字符不在byte的范围中 overflow溢出
	fmt.Printf("c3=%c c3对应的码值=%d\n",c3,c3)

	//可以直接给某个变量赋值一个数字,然后按照格式化输出时 `%c`,会输出该数字对应的unicode字符
	var c4 int  = 22269  //22369 -> '国'  120 -> 'x'
	fmt.Printf("c4=%c\n",c4)

	//字符类型是可以进行运算的,相当于一个整数,因为它都对应有`Unicode`码
	var n1 = 10 + 'a'
	fmt.Println("n1=",n1)
 }	
