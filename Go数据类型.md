
<h1> Go数据类型 <h1/>
#### 数字类型
Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码
 `int` 整型
` float32,float64` 浮点型 32位浮点型数，64位浮点型数
` uint8,uint16,uint32,uint64` 无符号整型 (8:0-255,16:0-65535,32:0-4294967295,64:0-18446744073709551615)
` int8,int16,int32,int64` 有符号整型  (-128 到 127，-32768 到 32767,-2147483648 到 2147483647,-9223372036854775808 到 9223372036854775807)
`complex64` 浮点型 32位实数和虚数
`complex128`浮点型 64位实数和虚数

#### 其他数字类型
`byte` 类似uint8
`rune` 类似int32
`uint` 32位或者64位
`int` 与uint一样大小
`uintptr` 无符号整型，用于存放一个指针

#### 字符串类型 
字符串就是一串固定长度的字符连接起来的字符序列。
Go 的字符串是由单个字节连接起来的。
Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

#### 布尔型

布尔型的值只可以是常量 true 或者 false。例子：var b bool = true。 注意：不能为其他类型


#### 派生类型

- 指针类型（Pointer）
- 数组类型
- 结构化类型 （struct)
- Channel 类型 (chan)
- 函数类型 (func)
- 切片类型 (slice)
- 接口类型 （interface）
- Map 类型 (map)


#### 注意
- 类型零值不是空值，而是某个变量被声明后的默认值，
- 一般情况下，值类型的默认值为0，
- 一般情况下，布尔默认值为false
- 一般情况下，字符串类型默认值为空字符串


#### 类型别名
通过 `type` 进行定义，如下
```
    type 整型 int32
    var  zhengxing 整型 = 1 
    fat.Print(zhengxing) // 打印会输出 1
```
注意：类型别名不能参与计算，相同别名的类型是可以的参与计算的，如下
```
    type 整型 int32
    var  zhengxing 整型 = 1 
    var jisuan int32 = 1
    fat.Print(zhengxing + jisuan) // 不能参与计算
    var zhengxing1 整型 = 1
    fat.Print(zhengxing + zhengxing1) // 是可以计算的
```


