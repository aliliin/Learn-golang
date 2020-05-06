package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 不加 * 是值传递，值传递是不能修改 里面的内容的，外面也收不到
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 不加 * 是值传递，值传递是不能修改 里面的内容的，加了* 是指针
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("是 nil 的调用")
		return
	}
	node.Value = value
}

/**
	值接收者 vs 指针接收者
  要改变内容必须使用指针接收者
 结构过大也考虑使用指针接收者
（建议）一致性：如果指针接收者，最好都是指针接收者
值接收者是 go 语言特有

*/

// 自定义工厂函数
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}
