package main

import (
	"fmt"
	//为了使用model.go的变量，需要先引入包
    "sheep.com.v8/model"
)
/*
 * @Description: 
 * @Author: yangyangSheep
 * @Date: 2022-04-22 17:53:42
 * @LastEditors: yangyangSheep
 * @LastEditTime: 2022-04-22 18:16:12
 */
func main()  {
	//引用的方法是使用 包名.变量名
	// fmt.Println(model.heroName)
	fmt.Println(model.HeroAge)
}