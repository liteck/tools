package tools

import (
	"sync"
)

type LQueues struct {
	running sync.Map
}

// 多任务按顺序执行
func NewQueues() *LQueues {
	q := new(LQueues)
	return q
}

func (q *LQueues) init0() {
	return
}

func (q *LQueues) AddTask(key string, task LTask) {
	var n *LQueue
	if v, ok := q.running.Load(key); !ok {
		// no data
		n = NewQueue(1)
	} else {
		n = v.(*LQueue)
	}
	n.AddTask(task)
	q.running.Store(key, n)
	printInfo("running queue's count is ")
	return
}

func (q *LQueues) Get(key string) (v interface{}, ok bool) {
	return q.running.Load(key)
}

func (q *LQueues) Close(key string) {
	printInfo("LQueues 准备关闭 " + key + " 的队列")
	var n *LQueue
	if v, ok := q.running.Load(key); !ok {
		printInfo("LQueues 关闭 " + key + " 的队列失败,没有找到这个队列,可能已经关闭了")
		return
	} else {
		n = v.(*LQueue)
	}
	n.Close()
	q.running.Delete(key)
	printInfo("LQueues 关闭 " + key + " 的队列成功")
}

func (q *LQueues) CloseAll() {
	printInfo("LQueues 准备退出,关闭所有的队列")
	q.running.Range(func(key, value interface{}) bool {
		n := value.(*LQueue)
		n.Close()
		q.running.Delete(key)
		return true
	})
	printInfo("LQueues 退出完毕")
}
