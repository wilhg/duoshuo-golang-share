golang talk
--

####1. 类型声明后置
不是什么深奥的事
在复杂函数声明的时候，会比类型前置更清晰可读

在变量声明的时候完全把type和value分开了
**c**:  int main(int argc, char \*argv[]) { /\* ... \*/ }

**go**: func main(argc int, argv *[]byte) int

####2. 多返回与错误处理
```go
a, b := "c", "c"    // right
a := b := "c"       // error
```

golang中，一般使用if来判断函数执行是否抛错，没有try-catch
```go
func a() err
if err := a(); err != nil {
    if true {
        return nil
    } else {
        return error.New("I'm error.")
    }
}
if err := a(); err != nil {

}

/////////
func b() (err, result) {
    if true {
        return nil, "Hello Golang"
    } else {
        return error.New("I'm error."), ""
    }
}
err, result := b()
if err != nil {
    /* ... */
}
// ignore the error
_, result := b()
```
####3. switch
非穿透型的，fallthrough

因为switch支持表达式，所以goblog中说，如果你的if需要判断三次以上，请使用switch
http://www.cnblogs.com/howDo/archive/2013/06/01/GoLang-Control.html

####4. 面向对象
```go
type Human struct {
    name string
    age int
    phone string
}

type Employee struct {
    Human //匿名字段
    company string
}

//Human定义method
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
        e.company, e.phone) //Yes you can split into 2 lines here.
}
```
> 通过变量首字母大小写区分public/private

https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.5.md
####5. interface
`interface{}`是一个空类型，golang中所有类型都实现了`interface{}`

传统定义：如果一只鸟是鸭子，那么它就是鸭子
鸭子类型：如果它会拍翅膀，会叫，会游泳，那么它是鸭子

```go
type Speaker interface {
     Speak() string
}
type Teacher struct {
    name string
}
func (t *Teacher) Speak() string {
    return "Attention!"
}
```
无需通过声明，只需直接实现接口方法。

接口可以内嵌，著名的例子：`io`包下的`ReadWriter`接口
```go
// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
    Write(p []byte) (n int, err error)
}

// io.ReadWriter
type ReadWriter interface {
    Reader
    Writer
}
```

####6. 并发控制
关键字：go, chan, select, close
> 不要通过共享来通信，而要通过通信来共享。

#```go```
**channel**:
```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})

ch := make(chan type, value) //channel缓冲区
```

channel通过操作符`<-`来接收和发送数据
```go
ch <- v    // 发送v到channel ch.
v := <-ch  // 从ch中接收数据，并赋值给v
```
**select** 用来对并发进行控制:
```go
for {
    select {
        case out <- n：
            // output or something
        case <-done：
            return
    }
}
```
**goroutine** 协程：
```go
func routine(data string, ch chan string) {
    // manipulate data
    ch <- data // ch is a point
}
```

[并发控制进阶](http://air.googol.im/2014/03/15/go-concurrency-patterns-pipelines-and-cancellation.html)
###what's more?
- reflection
- point
- stdlib
    - gowalker.org
    - go-search.org

golang.org很早就被GFW认证

gocode - 斯巴达不一定就是不好的
