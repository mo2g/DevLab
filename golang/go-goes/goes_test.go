package goes

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// Go Pool Test
//

func init() {
	println("using MAXPROC")
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
}

func TestAddTask(t *testing.T) {
	goes := NewGoesPoolDefault(10)
	goes.Start()
	defer goes.Shutdown()

	done := make(chan bool)
	scheduled := false

	goes.Add(func() {
		println("Called")
		scheduled = true
		close(done)
	})

	<-done
	if !scheduled {
		t.Error("Task is not be scheduled")
	}
}

func TestTask2(t *testing.T) {
	goes := NewGoesPoolDefault(100)
	goes.Start()
	defer goes.Shutdown()

	wg := new(sync.WaitGroup)

	TASKS := int(1000 * 100)
	for i := 0; i < TASKS; i++ {
		wg.Add(1)
		goes.Add(func() {
			wg.Done()
		})
	}

	timer := time.AfterFunc(time.Millisecond*100, func() {
		t.Error("Timeout 100ms")
	})

	wg.Wait()
	timer.Stop()

	t.Log("Send tasks:" + strconv.Itoa(TASKS))
}

func BenchmarkCpuNumWorkers(b *testing.B) {
	MakeBenchmarkWith(runtime.NumCPU(), b)
}

func Benchmark1KWorkers(b *testing.B) {
	MakeBenchmarkWith(1000, b)
}

func MakeBenchmarkWith(numWorkers int, b *testing.B) {
	goes := NewGoesPool(numWorkers, numWorkers)
	goes.Start()
	defer goes.Shutdown()

	count := 0
	for i := 0; i < b.N; i++ {
		goes.Add(func() {
			count++
		})
	}

}
