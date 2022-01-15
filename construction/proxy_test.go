package construction

import "testing"

func TestProxyDockerhub_Pull(t *testing.T) {
	p := &ProxyDockerhub{}
	p.Pull()
	realdockerhub := &RealDockerhub{
		Server: "hub.docker.xx",
	}
	if p.RealDockerhub != realdockerhub {
		t.Fatal("result not as expect")
	}
}
