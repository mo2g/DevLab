package goes

import (
	"testing"
	"time"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestWorker_Run(t *testing.T) {
	w := newWorker()
	defer close(w.tasks)

	idles := make(chan *worker, 1)
	go w.work(idles)

	done := make(chan bool)
	w.tasks <- func() {
		done <- true
	}

	time.AfterFunc(time.Millisecond, func() {
		t.Error("Timeout")
	})

	<-done

	select {
	case <-idles:
		t.Log("Test ok")

	default:
		t.Error("Notify idle not call")
	}
}

func BenchmarkWorker(b *testing.B) {
	w := newWorker()
	defer close(w.tasks)

	idles := make(chan *worker, b.N)
	go w.work(idles)

	count := 0
	for i := 0; i < b.N; i++ {
		w.tasks <- func() {
			count++
		}
	}
}
