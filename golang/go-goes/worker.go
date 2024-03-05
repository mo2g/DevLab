package goes

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type GoTask func() // 任务

// Worker包含一个任务列表和停止控制信号
type worker struct {
	tasks    chan GoTask   // 任务列表
	waitExit chan struct{} // 运行状态
}

// 启动内部协程，异步循环处理任务列表
func (slf *worker) work(idles chan<- *worker) {
	defer close(slf.waitExit)

	finishTask := func() {
		idles <- slf
	}

	for taskFunc := range slf.tasks {
		func() {
			// 完成一个任务后，将当前Worker恢复到空闲状态
			defer finishTask()

			taskFunc()
		}()
	}
}

func newWorker() *worker {
	return &worker{
		tasks:    make(chan GoTask, 1), // 1: for sync send task
		waitExit: make(chan struct{}),
	}
}
