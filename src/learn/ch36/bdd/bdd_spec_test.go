package bdd_test

// bdd 测试工具的使用
// github 地址 ： https://github.com/smartystreets/goconvey.git
import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		Convey("When the integer is incremented", func() {
			x++
			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
