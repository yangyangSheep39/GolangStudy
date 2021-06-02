# GolangStudy

This is Golang Study Note Directory

## 查看版本号

`go version`

## 查看env

`go env`

## 拓展链接

### web框架

[流行度排行](https://github.com/speedwheel/awesome-go-web-frameworks/blob/master/README.md#popularity)  
[beego](https://beego.me/products)  
[echo](https://echo.labstack.com/)  
[gin go社区文档](https://learnku.com/docs/gin-gonic/2019)  
[看云gin文档](https://www.kancloud.cn/shuangdeyu/gin_book/949418)  
[gin快速使用](https://www.jianshu.com/p/98965b3ff638)  
[gin-admin-vue](https://www.gin-vue-admin.com/docs/gorm)  
[官网](https://golang.org/dl/)  
[中文官网](https://go-zh.org/pkg/)  
[go常用库](https://github.com/jobbole/awesome-go-cn)  

### 拉不下来包

```Go
export GOPATH=/usr/local/go/bin
export GOPROXY=https://goproxy.io
export PATH=$PATH:$GOPATH
export GO111MODULE=on
```

### 开启module

```Go
go env -w GO111MODULE=on
```

### 设置代理

```Go
go env -w GOPROXY=https://goproxy.io,direct
```

### 微服务框架

> go-kit
> go micro
> [go-zero](https://go-zero.dev/cn/)  

### 常用网址

[go基础知识](http://www.topgoer.com)  
[go基础知识v1](https://chai2010.cn/advanced-go-programming-book/ch5-web/ch5-03-middleware.html)  
[go基础知识v2---micro](https://www.qfgolang.com/?special=go-microweifuwukuangjia&pid=2650)  
[菜鸟GO](https://www.runoob.com/go/go-tutorial.html)  
[文档和技术论坛](https://learnku.com/go)  
[gomod 详细使用](https://www.jianshu.com/p/760c97ff644c)  
[go官网库](https://golang.org/pkg/)  
[go官方pkg中文](https://studygolang.com/static/pkgdoc/main.html)  
[github对pgk的使用例子](https://github.com/astaxie/gopkg)  

### go mod使用

> ![go mod使用.png](./source/picture/go%20mod使用.jpg)  

## Table of Contents

### 1.Go的特性和快速体验

[Go语言的特点](./基础篇/1.Go的特性和快速体验/Go的特性和快速体验.md#Go语言的特点)  
[Google为什么要创造Go语言](./基础篇/1.Go的特性和快速体验/Go的特性和快速体验.md#Google为什么要创造Go语言)

### 2.Go的基本结构说明和HelloWorld程序开发快速入门

[Go的程序开发和基本结构说明](./基础篇/2.Go的基本结构说明和HelloWorld程序开发快速入门/Go的基本结构说明和HelloWorld程序开发快速入门.md#Go的程序开发和基本结构说明)  
[Golang的执行流程分析](./基础篇/2.Go的基本结构说明和HelloWorld程序开发快速入门/Go的基本结构说明和HelloWorld程序开发快速入门.md#Golang的执行流程分析)  
[编译和运行的说明](./基础篇/2.Go的基本结构说明和HelloWorld程序开发快速入门/Go的基本结构说明和HelloWorld程序开发快速入门.md#编译和运行的说明)  
[Go程序开发的注意事项](./基础篇/2.Go的基本结构说明和HelloWorld程序开发快速入门/Go的基本结构说明和HelloWorld程序开发快速入门.md#Go程序开发的注意事项)  
[Golang的常用转义字符escape-char](./基础篇/2.Go的基本结构说明和HelloWorld程序开发快速入门/Go的基本结构说明和HelloWorld程序开发快速入门.md#Golang的常用转义字符escape-char)  
[Golang的注释-comment](./基础篇/2.Go的基本结构说明和HelloWorld程序开发快速入门/Go的基本结构说明和HelloWorld程序开发快速入门.md#Golang的注释-comment)  
[Golang规范的代码风格](./基础篇/2.Go的基本结构说明和HelloWorld程序开发快速入门/Go的基本结构说明和HelloWorld程序开发快速入门.md#Golang规范的代码风格)  

### 3.DOS介绍以及常用命令

[DOS](./基础篇/3.DOS介绍以及常用命令/DOS.md#DOS)  

### 4.Golang变量

[概念](./基础篇/4.Golang变量/Golang变量.md#概念)  
[变量使用的基本步骤](./基础篇/4.Golang变量/Golang变量.md#变量使用的基本步骤)  
[变量的使用注意事项和细节](./基础篇/4.Golang变量/Golang变量.md#变量的使用注意事项和细节)  
[变量在代码运行中的变化](./基础篇/4.Golang变量/Golang变量.md#变量在代码运行中的变化)  
[变量的三种声明方式](./基础篇/4.Golang变量/Golang变量.md#变量的三种声明方式)  
[变量的声明,初始化和赋值](./基础篇/4.Golang变量/Golang变量.md#变量的声明,初始化和赋值)  
[程序中+号的使用](./基础篇/4.Golang变量/Golang变量.md#程序中+号的使用)  

### 5.Golang的数据类型

[整数类型](./基础篇/5.Golang的数据类型/Golang的数据类型.md#整数类型)  
[浮点类型](./基础篇/5.Golang的数据类型/Golang的数据类型.md#浮点类型)  
[字符类型](./基础篇/5.Golang的数据类型/Golang的数据类型.md#字符类型)  
[布尔类型](./基础篇/5.Golang的数据类型/Golang的数据类型.md#布尔类型)  
[字符串类型](./基础篇/5.Golang的数据类型/Golang的数据类型.md#字符串类型)  
