this is an example for how to write a simple orm

我们希望用自己动手写一个orm 的同时去学习builder模式
Builer().Table().Where().Limit().Do()
像上面这样

例如
查询
`insert into product (name, description, price) values (xx,xx,xx)`
就可以写成
```go
product := Product{
    Name: "testproduct",
}
Builer().Table("product").Do("insert", product)
```

删除
`delete from product where (name = testproduct)`
可以写成
```go
Builer().Table("product").Where("name", "=", "testproduct").Do("delete")
```

我们考虑增加删除语句的都是用k-v 键值对， k-v的只有map 跟struct
map的缺点是所有key 都是一个类型，这不符合我们的要求，因此使用struct
```go
type Product struct {
	Name string `sql: name`
}

p1 := Product{
	Name: "tt",
}
Newbuilder().Table("product").Insert(p1)
```





