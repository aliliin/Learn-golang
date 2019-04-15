package pipe_filter

import "testing"

// 架构模式。
// 面向对象的可复用的设计模式

// Pipe-Fiter 架构模式
/**
 * 非常适合数据处理及数据分析系统
 * Filter 封装数据处理的功能
 * 松耦合 ： Filter 只跟数据（格式） 耦合
 * Pipe 用于链接 Filter 传递数据或者在异步处理过程中缓冲数据流进程内部同步调用时，pipe 演变为数据在方法调用间传递
 */

//  Fiter 和组合模式
func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStaightPipeline("p1", spliter, converter, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}
