package factory

import (
	"reflect"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	b := NewBuilder()
	b.WithNamespace("default").WithFlatten(true).WithLabelSelectorParam("test")

	expect := &Builder{
		Name: "jane",
		Namespace: "default",
		Flatten: true,
		LabelSelectorParam: "test",
	}
	if !reflect.DeepEqual(b, expect) {
		t.Fatal("not as expect")
	}
}