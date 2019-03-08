
<h1> Go 基础语法 <h1/>

#### 关键字
Go 语言中保留的25个关键字如下表

| break | default |func | interface | select |
|---|---|---|---|---|
|case|defer|go|map|struct|
|chan|else|goto|package|switch|
|const|fallthrough|if|range|type|
|continue|for|import|return|var|

Go 语言 36 个预定义标识符

|append|bool|byte|cap|close|complex|complex64|complex128|uint16|
|---|---|---|---|---|---|---|---|---|
|copy|false|float32|float64|imag|int|int8|int16|uint32|
|int32| int64|iota|len|make|new|nil|panic|uint64|
|print| println|real|recover|string|true|uint|uint8|uintptr|



#### Package 包
* ` package` 是最基本的分发单位和工程管理中依赖关系的体现 
*  每个Go语言源代码文件开头都拥有一个`package`声明，表示源码文件所属代码包
*  要生成Go语言可以执行程序，必须要有`main`的`package`包，且必须在该包下有`main() `函数
*  同一个路径下只能存在一个`package` ，一个`package` 包可以拆成多个源文件组成

建议引入的`package`包的名称和当前属于的文件名相同

#### import 的用法及原理
-  `import `语句可以导入源代码文件所依赖的` package` 包
-  不得 导入源代码文件中没有用到的 `package` ，否则Go语言编译器会报编译错误
-  如果一个`main`导入其他包，包将按照顺序导入
-  如果导入的包中依赖其他包（A依赖B）,会首先导入B包，然后初始化B包中的常量和变量，最好如果B包中有` init `,会自动执行 `init() `初始化函数
-  所有包导入完成后才对main中的常量和变量进行初始化，然后执行`main`中的`init()` 函数(如过存在)，最后执行`main`函数
-  如果一个包被导入多次则该包之后导入一次。

import 语法：
第一种用法
```
	import "package1"
	import "package2"
	import "package3"
```
第二种用法
```
	import (
		"package1"
		"package2"
		"package3"
	)
```
#### import 特殊用法
- `import`  可以取别名，将导入的包命名为另一个容易记忆的别名 （别名写在 引入包的名称前面）
-  点（.）操作的含义是：点（.）标示的包导入后，调用该包中函数时可以省略前缀包名
-  下划线 （_）操作的含义是：导入该包，但不导入整个包，而是执行该包中的`init`函数,因此无法通过包名来调用包中的其他函数。使用下划线操作往往是为了注册包里的引擎，让外部可以方便的使用
- 

