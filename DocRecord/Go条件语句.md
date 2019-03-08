<h1>Go 条件语句</h1>

#### if 语句
- if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
- if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
- 你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。

```
	if 布尔表达式 {
	   /* 在布尔表达式为 true 时执行 */
	}
	
	number := 1	
	
	if number == 2 {
		fmt.Print("等于2")
	}else{
		fmt.Print("不等于2")
	}
	
	age := 28
	
	if age == 26{
		fmt.Print("age是26")
	}else if age < 27{
		fmt.Print("age是27")
	}else {
		fmt.Print("age大于27") // age大于27
	}

	// 嵌套if语句
	a := 100
	b := 10
	if a > b {
		fmt.Print("a大于等于b \n")
		if a > 20{
			fmt.Print("a大于于20 \n")
		}
	}else{
		fmt.Print("a小于b \n")
	}
```

#### switch 语句
- switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
- switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

```
	 // 基础语法
	switch var1 {
	    case val1:
	        ...
	    case val2:
	        ...
	    default:
	        ...
	}

	// 用来判断字符类型 
	var a interface{}
	a = 3.143
	
	switch a.(type) {
	case int:
		fmt.Print("类型为int\n")
	case string:
		fmt.Print("类型为字符串\n")
	default:
		fmt.Print("未知类型 \n") // 输出未知类型
	}
	
 // 判断成绩
	var grade string = "D"
	
	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" ) // 及格
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" );
	}
```

 select 语句暂时还不理解，后续再写。
