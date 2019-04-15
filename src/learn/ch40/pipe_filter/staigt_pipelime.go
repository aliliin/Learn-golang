package pipe_filter

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
func NewStaightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
