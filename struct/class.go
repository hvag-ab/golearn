package struct_
/**
Go语言通过首字母的大小写来控制访问权限。无论是方法，变量，常量或是自定义的变量类型，如果首字母大写，则可以被外部包访问，反之则不可以。
 */
import (
	"fmt"
)


// 模拟动物行为的接口
type IAnimal interface {
	Eat() // 描述吃的行为
}

// 动物 所有动物的父类
type Animal struct {
    Name string
}

// 动物去实现IAnimal中描述的吃的接口 
func (a *Animal) Eat() {
	a.Name = "hvag" // 传入指针 可以改变属性值
    fmt.Printf("%v is eating\n", a.Name)
}

// //值类型 不能改变实例里面的属性值
func (a Animal) Eat2() {
	a.Name = "hvagab" 
    fmt.Printf("%v is eating\n",a.Name)
}

// 不是严格上的继承 当子类调用sleep的时候 sleep调用了eat方法 是父类的eat方法 这里并不是继承上的子类的eat方法
func (a *Animal) sleep(){
	a.Eat()
	fmt.Printf("%v is sleep\n", a.Name)
}

// 动物的构造函数
func NewAnimal(name string) *Animal {
    return &Animal{
        Name: name,
    }
}

// 猫的结构体 组合了animal 相当于继承
type Cat struct {
    *Animal
}

// 实现猫的构造函数 初始化animal结构体
func NewCat(name string) *Cat {
    return &Cat{
        Animal: NewAnimal(name),
    }
}

// 猫结构体IAnimal的Eat方法 重复父类方法
func (cat *Cat) Eat() {
	// cat.Eat() //调用父类方法
	cat.Name = "cat"
    fmt.Printf("children %v is eating\n", cat.Name)
}

// 检查接口
func Check(animal IAnimal) {
    animal.Eat()
}





type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human
	phone string //human中也有phone字段
}

type student2 struct {
	hu human
	phone string //human中也有phone字段
}

//给基类增加walk方法
func (a human) walk() {
    fmt.Printf("%s are walking.. \n", a.name)
}

func stu() {
	// 重载字段，就近原则
	hvag := student{human{name: "hvag", age: 20, phone: "110"}, "119"}
	fmt.Println("hvag phone uumber is ", hvag.phone) //119
	hvag.walk()

	hv := student2{human{name: "hvag2", age: 30, phone: "130"}, "129"}
	fmt.Println("hvag phone uumber is ", hv.phone,hv.hu)
	hv.hu.walk()
}


// func main(){
// 	cat := newCat("dog")
// 	// cat.Eat()
// 	// cat.sleep()
// 	check(cat)
// 	// stu()
// }