package tools

import (
	"container/list"
	"sync"
	"time"
)

type LTask interface {
	Run()
}

type LTaskPrepare interface {
	Prepare()
}

type LTaskDone interface {
	Done()
}

type LTaskRunning struct {
	// 运行的起始时间
	StartTime int64
}

// 建立一个任务
type LQueue struct {
	// 退出信号.
	exit chan int
	// 时间
	ticker *time.Ticker
	// 存储的是所有的任务,此列表可以不停存储.不设上限.
	l *list.List
	// 最大数量
	maxTaskCount int
	wg           sync.WaitGroup
	mux          *sync.Mutex
	state        LTaskRunning
}

/*
创建一个新的任务队列
@maxTaskCount 最大执行数量
*/
func NewQueue(maxTaskCount int) *LQueue {
	q := new(LQueue)
	q.init0(maxTaskCount)
	return q
}

// 初始化.
func (q *LQueue) init0(maxTaskCount int) {
	q.exit = make(chan int)
	q.ticker = time.NewTicker(time.Second)
	q.maxTaskCount = maxTaskCount
	q.l = list.New()
	q.wg = sync.WaitGroup{}
	q.mux = new(sync.Mutex)
	q.state = LTaskRunning{}
	q.state.StartTime = time.Now().Unix()
	go q.taskListener()
	return
}

// 获取任务.
func (q *LQueue) taskListener() {
	for {
		select {
		case <-q.ticker.C:
			q.run()
			break
		case <-q.exit: //监听 信号
			q.ticker.Stop()
			// 删除所有的数据
			var next *list.Element
			for ele := q.l.Front(); ele != nil; ele = next {
				next = ele.Next()
				q.l.Remove(ele)
			}
			printInfo("LQueue 退出完毕")
			return
		}
	}
}

func (q *LQueue) run() {
	defer q.mux.Unlock()
	q.mux.Lock()

	idx := 0
	for {
		if idx >= q.maxTaskCount {
			break
		}
		if ele := q.l.Front(); ele == nil {
			// list is empty, break and wait
			break
		} else {
			idx++
			// got task and send to chan
			q.wg.Add(1)
			task := q.l.Remove(ele).(LTask)
			go q.runTask(task)
		}
	}
	q.wg.Wait()
	return
}

func (q *LQueue) runTask(task LTask) {
	defer q.wg.Done()
	task.Run()
	if v, ok := task.(LTaskDone); ok {
		v.Done()
	}
}

// 运行的状态
func (q *LQueue) RunningState() (r LTaskRunning) {
	r.StartTime = q.state.StartTime
	return r
}

// 增加一个任务
func (q *LQueue) AddTask(task LTask) {
	defer q.mux.Unlock()
	q.mux.Lock()
	q.l.PushBack(task)
	if v, ok := task.(LTaskPrepare); ok {
		v.Prepare()
	}
	return
}

func (q *LQueue) Close() {
	printInfo("LQueue 准备退出")
	q.exit <- 0
}
