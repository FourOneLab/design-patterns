# Functional Optional

Functional Options 这个编程模式。这是一个函数式编程的应用案例，编程技巧也很好，是目前 Go 语言中最流行的一种编程模式。

这种方式至少带来了 6 个好处：

1. 直觉式的编程；
2. 高度的可配置化；
3. 很容易维护和扩展；
4. 自文档；
5. 新来的人很容易上手；
6. 没有什么令人困惑的事（是 nil 还是空）。

什么时候使用 Functional Options 这种编程模式呢？编程过程中，常常会有封装 struct 的操作，而通常会针对该 struct 做一个 New func 的操作，为的就是方便 inject 相对应的 dependency 进去。那么就会碰到需要有 option 的时候，所谓 option 的时候，是指说有些字段设置是可以给 client 自由设定的，此外如果 client 没有设定，会有所谓的预设值。这种时候就很时候使用这种设计模式。

如下面的 Server 这个 struct。

```go
type Server struct {
    Addr string
    Port int
    Protocol string
    Timeout time.Duration
    MaxConns int
    TLS *tls.Config
}
```

## 使用初始化函数

*因为 Go 语言不支持重载函数，所以，你得用不同的函数名来应对不同的配置选项。*

```go
func NewDefaultServer(addr string, port int) (*Server, error) {
    return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
}

func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {    
    return &Server{addr, port, "tcp", 30 * time.Second, 100, tls}, nil
}

func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
    return &Server{addr, port, "tcp", timeout, 100, nil}, nil
}

func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls *tls.Config) (*Server, error) {
    return &Server{addr, port, "tcp", 30 * time.Second, maxconns, tls}, nil
}
```

初始化函数的问题在于，会随着字段越多 function 的 parameter 就越来越长，拆分为多个 function 时，随着参数的增多，function 的数量也越来越多。

## 使用配置对象

为了解决初始化函数带来的问题，可以增加一个Config struct。

```go
type Config struct {
    Protocol string
    Timeout time.Duration
    MaxConns int
    TLS *tls.Config
}

type Server struct {
    Addr string
    Port int
    Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
    // ...
}

// Using the default configuration
srv1, _ := NewServer("localhost", 9000, nil)
conf := ServerConfig{ Protocol:"tcp", Timeout: 60*time.Duration }
srv2, _ := NewServer("localhost", 9000, &conf)
```

使用 Config 的问题在于：

1. 有时候并不需要 Config，那么在 NewServer 的时候就要传入 nil 或者空结构体，并且需要判断是否是 `nil` 或是 Empty—— `Config{}`会让代码感觉不太干净
2. 在 NewServer 检查每一个 Config 的参数是不是为 zero value，如果是的话就当做没有给参数，然后使用预设值，如果将我们需要设置的就是 zero value，那么就会不生效了

## 使用Builder模式

这种方式不需要额外的 `Config` 类，使用链式的函数调用的方式来构造一个对象，只需要多加一个 `Builder` 类。

> 这个 Builder 类似乎有点多余，可以直接在 `Server` 上进行这样的 `Builder` 构造。但是，在处理错误的时候可能就有点麻烦，不如一个包装类更好一些。

```go
// 使用一个 builder 类来做包装
type ServerBuilder struct {
    Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
    sb.Server.Addr = addr
    sb.Server.Port = port

    // 其它代码设置其它成员的默认值

    return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
    sb.Server.Protocol = protocol
    return sb
}

func (sb *ServerBuilder) WithMaxConn(maxConn int) *ServerBuilder {
    sb.Server.MaxConns = maxConn
    return sb
}

func (sb *ServerBuilder) WithTimeOut( timeout time.Duration) *ServerBuilder {
    sb.Server.Timeout = timeout
    return sb
}

func (sb *ServerBuilder) WithTLS( tls *tls.Config) *ServerBuilder {
    sb.Server.TLS = tls
    return sb
}

func (sb *ServerBuilder) Build() (Server) {
    return sb.Server
}

// 新建对象
sb := ServerBuilder{}

// 开始配置
server, err := sb.Create("127.0.0.1", 8080).
    WithProtocol("udp").
    WithMaxConn(1024).
    WithTimeOut(30*time.Second).
    Build()
```

## Function options

这组代码传入一个参数，然后返回一个函数，返回的这个函数会设置自己的入参 `Server`，这个叫高阶函数。再定义一个 `NewServer()`函数，其中，有一个可变参数 `options` ，它可以传出多个上面的函数，然后使用一个 `for-loop` 来设置我们的 `Server` 对象。

```go
type Option func(*Server)

func Protocol(p string) Option {
    return func(s *Server) {
        s.Protocol = p
    }
}

func Timeout(timeout time.Duration) Option {
    return func(s *Server) {
        s.Timeout = timeout
    }   
}

func MaxConns(maxconns int) Option {
    return func(s *Server) {
        s.MaxConns = maxconns
    }
}

func TLS(tls *tls.Config) Option {
    return func(s *Server) {
        s.TLS = tls
    }
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
    // 初始化的时候带上默认值
    srv := Server{
        Addr: addr,
        Port: port,
        Protocol: "tcp",
        Timeout: 30 * time.Second,
        MaxConns: 1000,
        TLS: nil,
    }

    for _, option := range options {
        option(&srv)    
    }

    // ...

    return &srv, nil
}

s1, _ := NewServer("localhost", 1024)
s2, _ := NewServer("localhost", 2048, Protocol("udp"))
s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
```

这不但解决了使用 `Config` 对象方式的需要有一个 `config` 参数，但在不需要的时候，是放 `nil` 还是放 `Config{}`”的选择困难问题，也不需要引用一个 `Builder` 的控制对象，直接使用函数式编程，在代码阅读上也很优雅。

那我们就可以在 NewServer 这边先给 default value，然后通过 for loop 将每一个 options 对其 Server 做的参数进行设置，这样不仅可以针对他想要的参数进行设置，其他没设置到的参数也不需要特地给 zero value 或是默认值，完全封装在 NewServer 就可以了。

## Uber的加强版

[Uber go guide](https://github.com/uber-go/guide/blob/master/style.md#functional-options)

```go
type options struct {
  cache  bool
  logger *zap.Logger
}

type Option interface {
  apply(*options)
}

type cacheOption bool

func (c cacheOption) apply(opts *options) {
  opts.cache = bool(c)
}

func WithCache(c bool) Option {
  return cacheOption(c)
}

type loggerOption struct {
  Log *zap.Logger
}

func (l loggerOption) apply(opts *options) {
  opts.logger = l.Log
}

func WithLogger(log *zap.Logger) Option {
  return loggerOption{Log: log}
}

// Open creates a connection.
func Open(
  addr string,
  opts ...Option,
) (*Connection, error) {
  options := options{
    cache:  defaultCache,
    logger: zap.NewNop(),
  }

  for _, o := range opts {
    o.apply(&options)
  }

  // ...
}
```

这样的设计方式就又更细粒度，将所有 option 给值的方式又再进行了封装。通过设计一个 `Option interface`，里面用了 `apply function`，以及使用一个 `options struct` 将所有的 field 都放在这个 struct 里面，每一个 field 又会用另外一种 struct 或是 custom type 进行封装，并 implement apply function，最后再提供一个 `public function`。

这样的做法好处是可以针对每一个 option 作更细的 custom function 设计，例如选项的 description 为何？可以为每一个 option 再去 implement Stringer interface，之后提供 option 描述就可以调用 toString 了，设计上更加的方便！
