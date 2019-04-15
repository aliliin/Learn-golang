package pipe_filter

import (
	"errors"
	"strconv"
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

//  转为整型
var ToIntFiterWrongFormatError = errors.New("input data should be string")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFiterWrongFormatError
	}
	ret := []int{}

	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
