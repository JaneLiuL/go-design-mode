package factory

// 这是builder 模式
// 我印象比较深刻的是kubectl 跟descheduler 代码里面看到的builder模式
// 这个例子是模仿kubectl builder模式写出来的
type Builder struct {
	Name       string
	Namespace string
	LabelSelectorParam string
	Flatten bool
}
func NewBuilder() *Builder  {
	return &Builder{
		Name: "jane",
	}
}

func (b *Builder) WithNamespace(namespace string) *Builder {
	b.Namespace = namespace
	return b
}

func (b *Builder) WithLabelSelectorParam(LabelSelectorParam string) *Builder {
	b.LabelSelectorParam = LabelSelectorParam
	return b
}

func (b *Builder) WithFlatten(flatten bool) *Builder {
	b.Flatten = flatten
	return b
}