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

所以Do 其实是一个interface?




