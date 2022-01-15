package construction

import "fmt"

// 代理模式，是非常常见的
// 例如翻墙就是使用了代理，去火车站或者找黄牛买票，黄牛或者火车站也是一种代理

type DockerRegistery interface {
	Pull()
}

type RealDockerhub struct {
	Server string
}

func NewRealDockerhub(server string) *RealDockerhub {
	return &RealDockerhub{
		Server: "hub.docker.xx",
	}
}

func (d *RealDockerhub) Pull() {
	fmt.Println("pull image from docker hub")
}

type ProxyDockerhub struct {
	RealDockerhub *RealDockerhub
	ProxyURL      string
}

func NewProxyDockerhub() *ProxyDockerhub {
	return &ProxyDockerhub{
		ProxyURL: "test.xx.xx",
	}
}

func (p *ProxyDockerhub) Pull() {
	if p.RealDockerhub == nil {
		p.RealDockerhub = NewRealDockerhub(p.ProxyURL)
	}
	p.RealDockerhub.Pull()
}
