# 错误处理

## 错误的本质

Go 语言的函数支持多返回值，可以在返回接口把业务语义（业务返回值）和控制语义（出错返回值）区分开。返回 `result`、`err` 两个值，这样的好处：

1. 参数基本都是入参，返回接口把结果和错误分离，使得函数接口语义清晰
2. 错误参数如果要忽略，需要显示地忽略，使用 `_` 变量来忽略
3. 返回的 `error` 是一个接口，可以扩展自定义的错误处理

```go
// 处理多个不同类型的 error
if err != nil {
    switch err.(type) {
    case *json.SyntaxError:
        …
    case *ZeroDivisionError:
        …
    case *NullPointerError:
        …
    default:
        …
    }   
}
```

Go 语言的错误处理的方式，本质上是返回值检查，也兼顾了异常的一些好处如对错误的扩展。

## 资源清理

不同编程语言有不同的资源清理的编程模式：

- C 语言：使用 `goto fail` 在一个集中的地方清理
- C++：使用 `RAII` 模式，通过面向对象的代理模式，把需要清理的资源交给一个代理类，然后再析构函数来解决
- Java：在 `finally` 语句块中进行清理
- Go：使用 `defer` 关键词进行清理

## 包装错误

把执行的一些上下文都包装到另一个错误中，然后一起返回，同时提供一个方法能够获取原始的错误。

```go
type authorizationError struct {
    operation string
    err error // original error
}

func (e *authorizationError) Error() string {
    return fmt.Sprintf(“authorization failed during %s: %v”, e.operation, e.err)
}

type causer interface {
    Cause() error
}

func (e *authorizationError) Cause() error {
    return e.err
}
```

更通用的做法，使用[第三方库](http://github.com/pkg/errors)，`errors.Wrap()` 最好保证只调用一次，否则全是重复的调用栈。

```go
import "github.com/pkg/errors"
// 错误包装
if err != nil {
    return errors.Wrap(err, “read failed”)
}

// Cause接口
switch err := errors.Cause(err).(type) {
    case *MyError:
    // handle specifically
    default:
    // unknown error
}
```

一个 goroutine 处理多个 error，当编写有着重试策略的代码时，将多个 error 合并为一个会十分有用。

- HashiCorp 的 [go-multierror](https://github.com/hashicorp/go-multierror)，性能更好一些
- Uber 的 [multierr](https://github.com/uber-go/multierr)

一个 error 多个 goroutine，在操作多个 goroutine 来处理一个任务的时候，为了保证程序的正确性，正确地管理结果和错误汇总是有必要的。

我们希望使 goroutine 之间相互依赖，并且如果其中一个失败就取消他们。避免无谓工作的解决方案可以是加一个 context，并且，一旦一个 goroutine 失败，就会取消它，这恰好就是 [`errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup) 所提供的；当处理一组 goroutine 的时候，一个错误以及上下文传播。

这里一定要注意三点：

1. `context` 是谁传进来的？其它代码会不会用到，`cancel()` 只能执行一次
2. `g.Go()` 不带 `recover()` 的，为了程序的健壮，一定要自行 `recover()`
3. 并行的 goroutine 有一个错误就返回，而不是普通的 fan-out 请求后收集结果

## 常见问题

### error 与 panic

理论上 panic 只存在于 server 启动阶段，比如 config 文件解析失败，端口监听失败等，所有业务逻辑禁止主动 panic。因此，所有异步的 goroutine 都要用 recover 去兜底处理。

### 错误处理与资源释放

```go
type result struct {
    Err error
}

func worker(done chan error){
    err:= doSomething()
    result := new(result)
    if err != nil {
        result.Err = err
    }
    done <- result
}
```

在 main 函数中启动 goroutine 来执行 worker，结果通过 channel 返回，在 main 函数中千万不能关闭 `done` channel，因为无法确定 `doSomething()` 何时返回，写 closed channel 直接 panic。

因此，数据传输和退出控制，需要用单独的 channel 不能混, 一般用 context 取消异步 goroutine, 而不是直接 close channels。

### error 级连使用

```go
package main

import "fmt"

type myError struct {
 string
}

func (i *myError) Error() string {
 return i.string
}

func Call1() error {
 return nil
}

func Call2() *myError {
 return nil
}

func main() {
 err := Call1() // error nil
 if err != nil {
  fmt.Printf("call1 is not nil: %v\n", err)
 }

 err = Call2() // *main.myError nil
 if err != nil {
  fmt.Printf("call2 err is not nil: %v\n", err)
 }
}
```

当自定义错误类型的时候，就会遇到经典的 Nil is not nil 的问题，可以重新定义一个新的 error 变量或者统一 error 类型。

### 并发问题

go 内置类型除了 channel 大部分都是非线程安全的，error 也不例外，因此，不要并发对 error 赋值。

### error 要不要忽略

为了保证兼容，error 一定要处理，至少打印日志。

### errWriter

```go
type errWriter struct{
    w io.Writer
    err error
}

func (ew *errWriter) write(buf []byte){
    if ew.err!=nil{
        return
    }
    _,ew.err = ew.w.Write(buf)
}

ew := &errWriter{w: fd}
// 一起处理，不用每处理一个就检查一次 error
ew.write(po[a:b])
ew.write(po[c:d])
ew.write(po[e:f])

if ew.err!=nil{
    return ew.err
}
```
