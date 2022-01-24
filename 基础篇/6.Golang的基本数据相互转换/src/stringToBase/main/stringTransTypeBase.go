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
 * @LastEditTime: 2022-01-24 16:13:41
 */
func main() {
	var str string = "true"
	var b bool

	//strconv.ParseBool(str) 本身会返回两个值 func ParseBool(str string) return: (value bool, err error)
	//我只想获取到value bool,不想获取err 所以用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n",b,b)
}