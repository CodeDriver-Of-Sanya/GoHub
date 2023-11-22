package main

import "fmt"

// 定义一个老师的结构体将老师各个属性放到结构体中管理
type Teacher struct {
	//变量名字大写外界可以访问这个属性
	Name   string
	Age    int
	School string
}

func main() {
	//创建老师结构体实例
	var t1 Teacher
	fmt.Println(t1) //在未赋值时默认值：{0}
	t1.Name = "张家欢"
	t1.Age = 19
	t1.School = "三亚学院"
	fmt.Println(t1)
	fmt.Println(t1.Age + 10)

	//第二种创建的方式
	var t Teacher = Teacher{"钟天心", 19, "三亚学院"}
	fmt.Println(t)
	//t.Name = "钟天心"
	//t.Age = 19
	//t.School = "三亚学院"

	//第三种方法利用指针
	var a *Teacher = new(Teacher)
	//a 是指针 a指向的就是Teacher结构体的地址 ，应该给这个地址指向的对象的字段赋值：
	(*a).Name = "张家欢"
	(*a).Age = 19
	//为了符合程序员的编程习惯，go提供了简化的赋值方式
	a.School = "三亚学院" //go编译器对a.School转化（*a）.School = "三亚学院 "
	fmt.Println(*a)

}
