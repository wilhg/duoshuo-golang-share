package main

import (
  "fmt"
)
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // 为每个 cs 中的输入 channel 启动一个 output Goroutine。outpu 从 c 里复制数值直到 c 被关闭
    // 或者从 done 里接收到数值，之后 output 调用 wg.Done
    output := func(c <-chan int) {
        for n := range c {
            select {
            case out <- n:
            case <-done:
            }
        }
        wg.Done()
    }
}
func main() {
    // 构建 done channel，整个管道里分享 done，并在管道退出时关闭这个 channel
    // 以此通知所有 Goroutine 该推出了。
    done := make(chan struct{})
    defer close(done)

    in := gen(done, 2, 3)

    // 发布 sq 的工作到两个都从 in 里读取数据的 Goroutine
    c1 := sq(done, in)
    c2 := sq(done, in)

    // 处理来自 output 的第一个数值
    out := merge(done, c1, c2)
    fmt.Println(<-out) // 4 或者 9

    // done 会通过 defer 调用而关闭
}
