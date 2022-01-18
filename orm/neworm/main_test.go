package main

import (
	"testing"
)

func TestNewOrmEngine_GetTable(t *testing.T) {

}

func TestInsert(t *testing.T) {
	mq := mq{}
	builder, err := mq.Newbuilder("root", "123456aGVsbG93b3JsZAo=aGVsbG93b3JsZAo=", "127.0.0.1", "cmdb")
	if builder == nil {
		t.Fatal("new mysql connection fail")
	}
	if err != nil {
		t.Fatal(err)
	}
	//builder.Migrate(Product)
	product := Product{
		Name:        "test",
		Description: "testdescription",
		Price:       10,
	}
	result, err := builder.Table("product").Insert(product).Do()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
