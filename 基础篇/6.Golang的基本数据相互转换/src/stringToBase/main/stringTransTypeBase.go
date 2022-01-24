package main

import (
	"fmt"
	"strconv"
)

/*
 * @Description: string类型 转换 基本数据类型
 * @Author: yangyangSheep
 * @Date: 2021-07-27 16:33:58
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2022-01-24 16:26:38
 */
func main() {
	var str string = "true"
	var b bool

	//strconv.ParseBool(str) 本身会返回两个值 func ParseBool(str string) return: (value bool, err error)
	//我只想获取到value bool,不想获取err 所以用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n",b,b)

	//strconv.ParseInt()(s string, base int, bitSize int) return:(i int64, err error)
	//返回字符串表示的整数值，接受正负号。
    //base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
    //bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	var str2 string = "123456"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2,10,64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v \n",n1,n1)
	fmt.Printf("n2 type %T n2=%v \n",n2,n2)
}