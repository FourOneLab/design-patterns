Happen Before

Channel的发送和接收

`Channel`是 `goroutine` 之间同步的主要方式。

* 开始发送 —> 接收完成

* 开始接收 —> 发送完成

开始发送和开始接收的顺序不定，对于无缓冲`Channel`来说，接收完成先于发送完成。

Lock

`sync.Mutex`/`sync.RWMutex`，假设`n`在`m`之前，那么`n`的`Unlock`调用先于在`m`的`Lock`调用返回。

* 有写锁时，`RLock`的调用返回发生在`Unlock`之后

* 有读锁时，`RUnLock`的调用发生在`Lock`之前

Once

`once.Do(f)` 中f的返回先于任何其他`once.Do`的返回。

#Golang


#Golang

## Go的初始化

* 程序初始化运行在单个`goroutine`中，但是该`goroutine`可以创建其他并发运行的`goroutine`。

* 如果包`p`导入来包`q`，则`q`包`init`函数执行结束先于`p`包`init`函数的执行。

* Main函数的执行在所有`init`函数执行完成后。

## 面向对象编程

*面向对象编程方法的黄金法则— Program to an interface not an implementation*

抽象出一个通用的函数，它的参数是一个接口，这样实现了这个接口的结构体都可以作为参数传入，这就可以实现结构体与某个特定功能解耦。

## Slice

是一个结构体而不是数组。

```go

type slice struct {

array unsafe.Pointer // 指向存放数据的数组指针

len int // 长度有多大

cap int // 容量有多大

}

```

[image:4525B7E3-C6EF-4144-9EDC-E2EA5DFC4B73-32843-000199FB37367FE4/slice.png]

在结构体里面用数组指针的问题，数据会发生共享。

`append()`函数在`cap`不够用的时候，就会重新分配内存以扩大容量，如果够用，就有不会重新分配内存。

## Full Slice Expression

Full Slice Expression中，最后一个参数叫 ”Limited Capacity”，后续对该slice的`append()`操作会导致重新分配内存。

```go

dir1 := path[:sepIndex]

dir1 := path[:sepIndex:sepIndex] // Full Slice Expression

```

## 深度比较

比较两个结构体或者Map中的数据是否相同，需要使用深度比较，`reflect.DeepEqual()`。

对于`map`和`struct`来说，与字段的顺序无关，但是对于`array`或者`slice`来说，与元素的顺序有关。

## 接口完整性检查

```go

var _ interface = (*struct)(nil)

```

1. 将nil变量强制类型转换为某个结构体都指针类型

2. 定义一个 `_` 变量，该变量的类型是某个接口

3. 比较这两者是否相等，来判断结构体是否实现了该接口

## 性能提示

### 推荐做法

1. 把数字转换成字符串：`strconv.Itoa()` 比 `fmt.Sprintf()`快::1::倍。

2. 在`for-loop`中对某个`Slice`进行`append()`，先把容量扩容到位，可以避免内存重新分配以及系统默认2的N次方扩展后浪费内存。

3. 在`for-loop`中的某个固定正则表达式，使用`regexp.Compile()`编译正则表达式，性能提升::2::个数量级。

4. 使用`bytes.Buffer`或者`strings.Builder`来拼接字符串，性能会比`+`或者`+=`高::3~4::个数量级。

5. 尽可能使用并发的`goroutine`，使用`sync.WaitGroup`同步分片操作。

6. 避免在热代码中进行内存分配，会导致频繁的GC，使用`sync.Pool`来重用对象。

7. 使用`lock-free`的操作，避免使用`mutex`，尽可能使用` sync/Atomic`包。

8. 使用`I/O` 缓冲（`I/O`是一个非常慢的操作），使用`bufio.NewWrite()`和`bufio.Newreader()`有更高的性能。

9. 更高性能的协议，使用`protobuf`或`msgp`，`JSON`在序列化和反序列化的时候使用了反射。

10. 在`map`中，使用整型的`key`比字符串快（整型之间的比较比字符串快）。

### 避免做法

1. 尽可能避免把`string`转换为`[]byte`，这个转换会导致性能下降。


> 一种软件设计的方法，它的主要思想是把控制逻辑（如，开关）与业务逻辑（如，电器）分开，不要在业务逻辑里写控制逻辑，因为这样会让控制逻辑依赖于业务逻辑，而是反过来，让业务逻辑依赖控制逻辑。

* 控制逻辑：开关

* 业务逻辑：电器

不要在电器中实现开关，而是把开关抽象成一种协议，让电器依赖它。*这样可以降低程序复杂度，提高代码重用度。*

在Golang中如何实现反转控制？

## 控制依赖业务（业务中实现控制）

Undo就是控制。

```go

package main

import “errors”

type UndoableIntSet struct {

IntSet

functions []func()

}

func NewUndoableIntSet() UndoableIntSet {

return UndoableIntSet{NewIntSet(), nil}

}

func (s *UndoableIntSet) Add(x int) {

if !s.Contains(x) {

s.data[x] = true

s.functions = append(s.functions, func() { s.Delete(x) })

} else {

s.functions = append(s.functions, nil)

}

}

func (s *UndoableIntSet) Delete(x int) {

if s.Contains(x) {

delete(s.data, x)

s.functions = append(s.functions, func() { s.Add(x) })

} else {

s.functions = append(s.functions, nil)

}

}

func (s *UndoableIntSet) Undo() error {

if len(s.functions) == 0 {

return errors.New(“no functions to undo”)

}

index := len(s.functions) - 1

if function := s.functions[index]; function != nil {

function()

s.functions[index] = nil

}

s.functions = s.functions[:index]

return nil

}

```

## 业务依赖控制

```go

package main

import "errors"

type IntSet struct {

data map[int]bool

undo Undo

}

func NewIntSet() IntSet {

return IntSet{data: make(map[int]bool)}

}

func (s *IntSet) Undo() error {

return s.undo.Undo()

}

func (s *IntSet) Add(x int) {

if !s.Contains(x) {

s.data[x] = true

s.undo.Add(func() { s.Delete(x) })

} else {

s.undo.Add(nil)

}

}

func (s *IntSet) Delete(x int) {

if s.Contains(x) {

delete(s.data, x)

s.undo.Add(func() { s.Add(x) })

} else {

s.undo.Add(nil)

}

}

func (s *IntSet) Contains(x int) bool {

return s.data[x]

}

type Undo []func()

func (u *Undo) Add(function func()) {

*u = append(*u, function)

}

func (u *Undo) Undo() error {

functions := *u

if len(functions) == 0 {

return errors.New("no functions to undo")

}

index := len(functions)

if function := functions[index]; function != nil {

function()

functions[index] = nil

}

*u = functions[:index]

return nil

}

```

## Map Reduce Filter 典型的控制与业务分离

```go

func MapStrToInt(arr []string, fn func(s string) int) []int {

var res []int

for _, s := range arr {

res = append(res, fn(s))

}

return res

}

func Reduce(arr []string, fn func(s string) int) int {

var res int

for _, s := range arr {

res += fn(s)

}

return res

}

func Filter(arr []int, fn func(n int) bool) []int {

var res []int

for _, i := range arr {

if fn(i) {

res = append(res, i)

}

}

return res

}

```

真正的业务逻辑由传入的数据和函数来确定。