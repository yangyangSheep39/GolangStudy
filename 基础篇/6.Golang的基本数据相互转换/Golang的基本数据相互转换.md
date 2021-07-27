# Golang的基本数据类型相互转换

## 介绍_base

Golang和Java/C不同,Golang在不同类型的变量之间赋值时需要**显式转换**,也就是说Golang中数据类型**不能自动转换**

### 基本语法

表达式 **T(v)** 将**值v**转换为**类型T**  
T:就是数据类型 比如int32,int64,float32等等  
v:就是需要转换的变量  

#### 案例演示

> [Golang的基本数据相互转换代码示例](./src/transBaseType/main/TransType.go)  

### 细节问题

1. Golang中,数据类型的转换是可以从 **表示范围小 -> 表示范围大**,也可以 **表示范围大 -> 表示范围小**  
2. 被转换的是**变量存储的数据(即值)**,变量本身的数据类型并没有变化  
3. 在转换中,比如将**int64转换成int8[-128~+127]**,编译时不会报错,只是转换的结果是**按溢出处理**,和我们希望的结果不一样,因此在转换时需要考虑范围的问题   

# 基本数据类型和string类型的转换

## 介绍_string

在程序开发中,我们经常需要将基本数据类型转换成string类型,或者将string类型转换成基本数据类型

### 基本数据类型转string类型

> 方式1: `fmt.Sprintf("%参数",表达式)` **【比较灵活】**  

1. 参数需要和表达式的数据类型相匹配  
2. `fmt.Sprintf()` 会返回转换后的字符串
3. 案例演示 [基本数据类型转string类型](./src/baseToString/main/baseTransTypeString.go)  

> 方式2: 使用 `strconv` 包的函数  

```Go
func Itoa
func Itoa(i int) string
Itoa是FormatInt(i, 10) 的简写。

func FormatBool
func FormatBool(b bool) string
根据b的值返回"true"或"false"。

func FormatInt
func FormatInt(i int64, base int) string
返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。

func FormatUint
func FormatUint(i uint64, base int) string
是FormatInt的无符号整数版本。

func FormatFloat
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
函数将浮点数表示为字符串并返回。

bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。

fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。

prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
```

### string类型转基本数据类型

> 使用 `strconv` 包的函数

1. 函数说明  

    ```Go
    func ParseBool
    func ParseBool(str string) (value bool, err error)
    返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。

    func ParseInt
    func ParseInt(s string, base int, bitSize int) (i int64, err error)
    返回字符串表示的整数值，接受正负号。

    base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；

    bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

    func ParseUint
    func ParseUint(s string, base int, bitSize int) (n uint64, err error)
    ParseUint类似ParseInt但不接受正负号，用于无符号整型。

    func ParseFloat
    func ParseFloat(s string, bitSize int) (f float64, err error)
    解析一个表示浮点数的字符串并返回其值。

    如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
    ```

2. 案例演示  [string类型转基本数据类型](./src/stringToBase/main/stringTransTypeBase.go)  
