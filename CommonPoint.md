# 基础知识细节问题汇总

## 包引入的方式

```Go
import "fmt"
import "unsafe"
//上述的引入包方式不够简洁,可以通过
import (
    _ "fmt" //如果我们没有是用到一个包 在前面加上一个下划线_ 表示忽略即可
    _ "unsafe"
)
```
