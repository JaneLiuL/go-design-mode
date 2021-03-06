package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type Builder struct {
	DB        *sql.DB
	TableName string
	Prepare   string
	AllExec   []interface{}
	Limit     string
	ExecWhere []string
}

type Product struct {
	Name        string `sql: "name"`
	Description string `sql: "description"`
	Price       int    `sql: "price"`
}

type mq struct {
}

func (mq mq) Newbuilder(username, password, addr, dbname string) (*Builder, error) {
	// also can enhance with builder mode
	//"root:123456aGVsbG93b3JsZAo=aGVsbG93b3JsZAo=@/cmdb"
	dsn := username + ":" + password + "@/" + dbname
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &Builder{
		DB: db,
	}, nil
}

// 设置表名
func (builder *Builder) Table(name string) *Builder {
	builder.TableName = name
	return builder
}

// 读取表名
func (builder *Builder) GetTable() string {
	return builder.TableName
}

func (builder *Builder) Where(wherecondition []string) *Builder {
	builder.ExecWhere = wherecondition
	return builder
}

func (builder *Builder) insertData(product interface{}, insertType string) *Builder {
	// we separte the product with the insert into xxx statement
	p := reflect.TypeOf(product)
	v := reflect.ValueOf(product)

	var fieldName []string
	var placeholder []string

	//p.NumField()可以获取到这个结构体有多少个字段用于 for 循环
	for i := 0; i < p.NumField(); i++ {
		// 小写开头，无法反射则跳过
		if !v.Field(i).CanInterface() {
			continue
		}

		// 解析tag 找出真实的sql字段名
		// p.Field(i).Tag.Get("sql")可以获取到包含sql:"xxx"的 tag 的值
		sqlTag := p.Field(i).Tag.Get("sql")
		if sqlTag != "" {
			//跳过自增字段
			if strings.Contains(strings.ToLower(sqlTag), "auto_increment") {
				continue
			} else {
				fieldName = append(fieldName, strings.Split(sqlTag, ",")[0])
				placeholder = append(placeholder, "?")
			}
		} else {
			fieldName = append(fieldName, p.Field(i).Name)
			placeholder = append(placeholder, "?")
		}
		// 把所有的值都放入到了builder.AllExec这个属性里面，之所以它用interface类型，是因为，结构体里面的值的类型是多变的，有可能是 int 型，也可能是 string 类型。
		builder.AllExec = append(builder.AllExec, v.Field(i).Interface())
	}
	//	拼接成这样 "insert into product (name, description, price) values (?,?,?)"
	builder.Prepare = insertType + " into " + builder.GetTable() + " (" + strings.Join(fieldName, ",") + ") values(" + strings.Join(placeholder, ",") + ")"
	//	separate value and put it into Exec
	return builder
}

func (builder *Builder) Insert(data interface{}) *Builder {
	return builder.insertData(data, "insert")
}

func (builder *Builder) Replace(data interface{}) *Builder {
	return builder.insertData(data, "replace")
}

func (builder *Builder) Migrate(data interface{}) {

}

func (builder *Builder) Do() (sql.Result, error) {
	fmt.Println(builder.Prepare)
	var statement *sql.Stmt
	statement, err := builder.DB.Prepare(builder.Prepare)
	if err != nil {
		return nil, err
	}
	//statement.Exec(e.AllExec...)，三个点的操作符，它能将我们传入的切片，全部拆开，一个的一个传入
	result, err := statement.Exec(builder.AllExec...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	mq := mq{}
	builder, err := mq.Newbuilder("root", "123456aGVsbG93b3JsZAo=aGVsbG93b3JsZAo=", "127.0.0.1", "cmdb")
	if builder == nil {
		panic("init fail")
	}
	if err != nil {
		fmt.Println(err)
	}
	//builder.Migrate(Product)
	product := Product{
		Name:        "test",
		Description: "testdescription",
		Price:       10,
	}
	builder.Table("product").Insert(product).Do()
}
