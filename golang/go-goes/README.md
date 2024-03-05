# Go Goes

> 50行代码实现非常简洁的Goroutines协程池 - Go

## Install

> go get -u github.com/yoojia/go-goes

## Dep

```toml
[[constraint]]
  name = "github.com/yoojia/go-goes"
  version = "0.1.0"
```

## Usage

三个步骤：

1. 创建GoesPool；
2. 启动协程池；
3. 通过Add方法添加任务函数；

```go
// 指定协程池创建100个工作协程
pool := goes.NewGoesPoolDefault(100)
// 启动
pool.Start()

// 停止协程池，阻塞等待所有任务函数执行完毕后返回。
defer pool.Shutdown()

wg := new(sync.WaitGroup)

TASKS := int(1000 * 100)
for i := 0; i < TASKS; i++ {
    wg.Add(1)
    // 通过 Add 方法，向协程池添加等待调度的任务函数
    pool.Add(func() {
        wg.Done()
    })
}

wg.Wait()

```

## Note

需要注意的问题：

1. 必须调用Start方法来启动协程池，否则Add的任务不会被执行；
2. Shutdown必须在调用Start之后调用，否则将**一直阻塞不会返回**；
3. Add方法，如果协程池已满，将阻塞，直到协程池有空闲Worker协程；

## Projects

本项目同时在以下两个代码托管平台更新维护：

- Microsoft.GitHub: [github.com/yoojia/go-goes](https://github.com/yoojia/go-goes)
- Gitee: [gitee.com/yoojia/go-goes](https://gitee.com/yoojia/go-goes)

## License

This package is licensed under Apache License 2.0. See [LICENSE](./LICENSE) for details.