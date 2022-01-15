package construction

// 适配器模式是把一种接口适配另外一种接口
// 例如，我们本来的output是json 格式以http协议给一个组件，现在我们加一个adapter让它同时支持 xml 格式以 RPC 协议给另外一个组件

type Output struct {
	Username string
	Age int
	Address string
}

