package interface_abstract

import "net/rpc"

// AuthenticationFilter 和 RateLimitFilter 是接口的两个实现类，分别实现了对 RPC 请求鉴权和限流的过滤功能。

// Filter 接口
type Filter interface {
	doFilter(req rpc.Request) error
}

// AuthenticationFilter 接口实现类：鉴权过滤器
type AuthenticationFilter struct{}

func (a AuthenticationFilter) doFilter(req rpc.Request) error {
	panic("implement me")
	// TODO: 鉴权逻辑
}

// RateLimitFilter 接口实现类：限流过滤器
type RateLimitFilter struct{}

func (r RateLimitFilter) doFilter(req rpc.Request) error {
	panic("implement me")
	// TODO: 限流逻辑
}

// Application 过滤器使用类
type Application struct {
	filters []Filter
}

func (a Application) HandleRPCRequest(req rpc.Request) error {
	for _, filter := range a.filters {
		if err := filter.doFilter(req); err != nil {
			return err
		}
	}
	return nil
}
