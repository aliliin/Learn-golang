package pipe_filter

import (
	"errors"
	"strings"
)

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

var SplitFiterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) // 检查数据格式/类型，是否可以处理
	if !ok {
		return nil, SplitFiterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
