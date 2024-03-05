package goes

//
// Author: 陈永佳 yoojiachen@gmail.com
//

// Goes内部维护一组Worker。
// 如果接收到外部任务，向空闲Worker发送任务
type GoesPool struct {
	workers       []*worker     // 所有Worker
	idles         chan *worker  // 空闲
	tasksQueue    chan GoTask   // 待调度任务列表
	waitCompleted chan struct{} // 终止状态
}

// Start 启动一个内部循环子协程来处理调度任务。
func (slf *GoesPool) Start() {
	go func() {
		defer close(slf.waitCompleted)

		for task := range slf.tasksQueue {
			// 取空闲Worker，发任务给它处理
			(<-slf.idles).tasks <- task
		}
		// 停止内部循环后，关闭Workers
		for _, worker := range slf.workers {
			close(worker.tasks)
			<-worker.waitExit
		}
	}()
}

// Shutdown 关闭GoesPool，阻塞等待所有Worker协程完成退出后，此函数才返回。
func (slf *GoesPool) Shutdown() {
	close(slf.tasksQueue)
	// 等待所有任务完成
	<-slf.waitCompleted
}

// Post 添加需要调度的任务。
// 当GoesPool有空闲Worker协程时，任务将被调度执行。否则，此函数将被阻塞，直到有空闲的Worker协程。
// 注意：Add 方法不能在Shutdown后调用，否则将会引发Panic。
func (slf *GoesPool) Post(task GoTask) {
	slf.tasksQueue <- task
}

func (slf *GoesPool) Add(task GoTask) {
	slf.Post(task)
}

// NewGoesPool 创建一个GoesPool新对象指针，指定内部Worker协程的数量。
func NewGoesPoolDefault(numWorkers int) *GoesPool {
	return NewGoesPool(numWorkers, max(1, numWorkers/2))
}

// NewGoesPool 创建一个GoesPool新对象指针，指定内部Worker协程的数量及Task任务列表容量。
func NewGoesPool(numWorkers int, taskQueueSize int) *GoesPool {
	numWorkers = max(1, numWorkers)
	taskQueueSize = max(1, taskQueueSize)
	goes := &GoesPool{
		workers:       make([]*worker, numWorkers),
		idles:         make(chan *worker, numWorkers),
		tasksQueue:    make(chan GoTask, taskQueueSize),
		waitCompleted: make(chan struct{}),
	}

	// 初始化Worker列表
	for i := 0; i < numWorkers; i++ {
		worker := newWorker()
		go worker.work(goes.idles)
		goes.workers[i] = worker
		goes.idles <- worker
	}

	return goes
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
