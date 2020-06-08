package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"//横杠表示加载包中的init函数
	"fmt"

	"log"
)

//pk 主键，autoincr自增 notnull null unique index index(indexname)联合索引 unique(uniquename)联合唯一 default(0) json
// -这个Field将不进行字段映射；->这个Field将只写入到数据库而不从数据库读取 <-这个Field将只从数据库读取，而不写入到数据库
//created 这个Field将在Insert时自动赋值为当前时间 updated这个Field将在Insert或Update时自动赋值为当前时间
//支持是数据库中所有类型 比如varchar char time datetime  首字母大写
type Account struct {//通过结构体tag 来定义表字段类型 tag就是``反引号里面的 deleted这个Field将在Delete时设置为当前时间，并且当前记录不删除
    Id      int64  `xorm:"int pk autoincr"` //int, int8, int16, int32, uint, uint8, uint16, uint32 都对应int int64 uint64对应BigInt
	Name    string `xorm:"Varchar(64) unique notnull 'user_name'"`//默认映射成varchar（255） 单引号扩起来的表示映射成数据库中哪一个字段
	
    Balance float64 `xorm:"Double"`//默认映射成double float 默认映射成float bool映射成1 or 0
    Version int `xorm:"version"` // 乐观锁
}

type auth struct {
	Username string 
	Password string 
}

var x *xorm.Engine

func main(){
	x,_:=xorm.NewEngine("mysql", "root:hvag@tcp(127.0.0.1:3306)/go?charset=utf8")
	x.ShowSQL(true)
	defer x.Close()
	// if err = x.Sync2(new(Account)); err != nil {//迁移表格
    //     log.Fatalf("Fail to sync database: %v\n", err)
	// }

	// a := &Account{Name:"hv", Balance:3.14}
	// DeleteOne(x,a)

	// acc := make([]*Account, 2)
	// acc[0] = &Account{Name:"abc",Balance:12.3}
	// acc[1] = &Account{Name:"xyz",Balance:1.634}
	// InsertMany(x,acc)

	// a := &Account{}
	a := new(Account)
	QueryGet(x,a)

	// QueryFind(x,a)

}

func InsertOne(x *xorm.Engine,a *Account){
	row, err := x.Insert(a)
	log.Fatalf("Fail to insert database: %v,%v\n", row,err)
	return 
}

func InsertMany(x *xorm.Engine,a []*Account){
	row, err := x.Insert(a)
	log.Fatalf("Fail to insert database: %v,%v\n", row,err)
	return 
}

func DeleteOne(x *xorm.Engine,a *Account){
	row, err := x.Delete(a)
	log.Fatalf("Fail to delete database: %v,%v\n", row,err)
	return 
}

func QueryGet(x *xorm.Engine,a interface{}){
	// has, err := x.Id(1).Get(a) //has表示是否存在 a表示查询哪一个表传递一个结构体地址&Account{} 这是根据id查询
	// 

	// has, err := x.Where("user_name=?", "hvag").Get(a)
	// has, err := x.Where("user_name=?", "hvag").Exist(a)

	// ma := map[string]interface{}{"user_name":"hvag"}//切记 一定是这种形式的字典才能接受
	// has, err := x.Where(ma).Get(a)


	b := &auth{Username:"hvag"}//根据结构体查询
	fmt.Println(b)
	// has, err := x.Get(b)//// 执行结束后，Account会被赋值为数据库中Id为1的实体
	has,err := x.Exist(b)	// Account中仍然是初始声明的Account，不做改变 如果仅仅是查询是否存在就选exist 效率更高
	log.Fatalf("Fail to query database: %v,%v\n", has,err,b)
	return
}


func QueryFind(x *xorm.Engine,a interface{}){
	// acc := make([]Account, 0)
	// err := x.Find(&acc)
	// log.Fatalf("Fail to query database: %v,%v\n", acc,err)
	// pacc := make([]*Account, 0)
	// err := x.Find(&pacc)
	// log.Fatalf("Fail to query database: %v,%v\n", pacc,err)


	// accmap := make(map[int64]Account)
	// err := x.Find(&accmap)//查询出来的值会全部映射到accmap中 返回err
	// log.Fatalf("Fail to query database: %v,%v\n", accmap,err)

	// paccmap := make(map[int64]*Account)
	// err := x.Find(&paccmap)

	// accs := make([]Account, 0)//where配合
	// err := x.Where("age > ? or name = ?", 30, "xlw").Limit(20, 10).Find(&accs)

	// var ints []int64 //查询单个字段返回切片
	// err := x.Table("account").Cols("id").Find(&ints) //Cols是查询特定字段 Omit是排除特定字段
	// log.Fatalf("Fail to query database: %v,%v\n", ints,err)

	// a := new(Account)
	// //返回满足id>1的Account的记录条数
	total, err := x.Where("id >?", 1).Count(a)
	log.Fatalf("Fail to query database: %v,%v\n", total,err)
	// ／／返回Account所有记录条数
	// total,err = x.Count(a)

	return
}

func UpdateData(x *xorm.Engine){//乐观所
	var acc Account
	x.Id(1).Get(&acc)
	v := acc.Version //每次更新时查看version是否变化 变化了说明已经更新了 
	acc2 := Account{Name:"ia",Balance:3.42,Version:v}
	// SELECT * FROM user WHERE id = ?
	x.Id(1).Update(&acc2)
// UPDATE user SET ..., version = version + 1 WHERE id = ? AND version = ?
}



/*

1、查询一个string类型的sql，返回[]map[string][]byte类型的切片（查询）

results, err := x.Query("select * from user")
　　2、执行一个string的sql，返回结果影响行数（增删改）

affected, err := x.Exec("update user set .... where ...")

sql := "update `userinfo` set username=? where id=?"
res, err := engine.Exec(sql, "xiaolun", 1) 


err := x.Iterate(...)
// SELECT * FROM user

raws, err := x.Raws(...)
// SELECT * FROM user
bean := new(Struct)
for raws.Next() {
    err = raws.Scan(bean)
}

affected, err := engine.Update(&user)
// UPDATE user SET 

affected, err := engine.Where(...).Delete(&user)
// DELETE FROM user Where 

engine.Id(1).Get(&user) // for single primary key
// SELECT * FROM user WHERE id = 1
engine.Id(core.PK{1, 2}).Get(&user) // for composite primary keys
// SELECT * FROM user WHERE id1 = 1 AND id2 = 2
engine.In("id", 1, 2, 3).Find(&users)
// SELECT * FROM user WHERE id IN (1, 2, 3)
engine.In("id", []int{1, 2, 3})
// SELECT * FROM user WHERE id IN (1, 2, 3)

engine.Where().And().Or().Find()
engine.Id(1).And(" user_name = ?",'davie").Get(&user)
// SELECT * FROM user WHERE (.. AND ..) OR ...
engine.Asc().Desc().Find()
// SELECT * FROM user ORDER BY .. ASC, .. DESC
engine.OrderBy().Find()
// SELECT * FROM user ORDER BY ..

engine.Limit().Find()
// SELECT * FROM user LIMIT .. OFFSET ..
engine.Top(5).Find()
// SELECT TOP 5 * FROM user // for mssql
// SELECT * FROM user LIMIT .. OFFSET 0 //for other databases



engine.GroupBy("name").Having("name='xlw'").Find()
//SELECT * FROM user GROUP BY name HAVING name='xlw'
engine.Join("LEFT", "userdetail", "user.id=userdetail.id").Find() //left right inner 大写
//SELECT * FROM user LEFT JOIN userdetail ON user.id=userdetail.id

指定某个表
x.Table("account")

别名
engine.Alias("o").Where("o.name = ?", name).Get(&order)

engine.Select("a.*, (select name from b limit 1) as name").Find(&beans)

engine.Where("column like ?", "%"+char+"%").Find(&beans) //like操作

engine.SQL("select * from table").Find(&beans)

engine.Select("a.*, (select name from b limit 1) as name").Get(&bean)

engine.Where("a = ? AND b = ?", 1, 2).Find(&beans)

engine.Where(builder.Eq{"a":1, "b": 2}).Find(&beans)

engine.Where(builder.Eq{"a":1}.Or(builder.Eq{"b": 2})).Find(&beans)

// select from table where column in (1,2,3)
engine.In("cloumn", 1, 2, 3).Find()

// select from table where column in (1,2,3)
engine.In("column", []int{1, 2, 3}).Find()

// select from table where column in (select column from table2 where a = 1)
engine.In("column", builder.Select("column").From("table2").Where(builder.Eq{"a":1})).Find()

engine.Cols("age", "name").Get(&usr)
// SELECT age, name FROM user limit 1
engine.Cols("age", "name").Find(&users)
// SELECT age, name FROM user
engine.Cols("age", "name").Update(&user)
// UPDATE user SET age=? AND name=?

engine.AllCols().Id(1).Update(&user)
// UPDATE user SET name = ?, age =?, gender =? WHERE id = 1f

// 例1：
engine.Omit("age", "gender").Update(&user)
// UPDATE user SET name = ? AND department = ?
// 例2：
engine.Omit("age, gender").Insert(&user)
// INSERT INTO user (name) values (?) // 这样的话age和gender会给默认值
// 例3：
engine.Omit("age", "gender").Find(&users)
// SELECT name FROM user //只select除age和gender字段的其它字段

engine.Distinct("age", "department").Find(&users)
// SELECT DISTINCT age, department FROM user

禁用自动根据结构体中的值来生成条件

engine.Where("name = ?", "lunny").Get(&User{Id:1})
// SELECT * FROM user where name='lunny' AND id = 1 LIMIT 1
engine.Where("name = ?", "lunny").NoAutoCondition().Get(&User{Id:1})
// SELECT * FROM user where name='lunny' LIMIT 1

UseBool(…string)
当从一个struct来生成查询条件或更新字段时，xorm会判断struct的field是否为0,“”,nil，如果为以上则不当做查询条件或者更新内容。
因为bool类型只有true和false两种值，因此默认所有bool类型不会作为查询条件或者更新字段。如果可以使用此方法，
如果默认不传参数，则所有的bool字段都将会被使用，如果参数不为空，则参数中指定的为字段名，则这些字段对应的bool值将被使用。
*/



