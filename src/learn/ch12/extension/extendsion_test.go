package extension_test

// go 继承
import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speack() {
	fmt.Print("...")
}
func (p *Pet) SpeackTo(host string) {
	p.Speack()
	fmt.Println(" ", host)

}

type Dog struct {
	p *Pet
}

func (d *Dog) Speack() {
	// d.p.Speack()
	// 重载此方法
	fmt.Println("汪汪汪")
}
func (d *Dog) SpeackTo(host string) {
	// d.p.SpeackTo(host)
	d.Speack()
	fmt.Println(" ", host)

}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeackTo("狗")
}

// 匿名内嵌套类型  不能继承 不支持方法重载
type Cat struct {
	Pet
}

func (c *Cat) Speack() {
	fmt.Println("miaomiaomaio")
}
func TestCat(t *testing.T) {
	cat := new(Cat)
	cat.SpeackTo("猫")

}
