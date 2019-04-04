package obj_pool

// 对象池，用于数据库连接啊网络连接啊
import (
	"errors"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct { // 用于可缓冲重用对象
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	ObjPool := ObjPool{}
	ObjPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		ObjPool.bufChan <- &ReusableObj{}
	}
	return &ObjPool
}

// 获取
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time Out")
	}
}

// 放回
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

// 换成空接口 可以放置任何对象 不建议这样子使用
func (p *ObjPool) GetAnyObj(timeout time.Duration) (obj interface{}, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time Out")
	}
}
