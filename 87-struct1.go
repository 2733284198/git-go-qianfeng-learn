package main

import "fmt"

func main() {
	fmt.Println("\n\n ===>72 fun <====")

	//	 使用
	var p Person
	p.name = "max"
	p.age = 20

	// student
	var stu Student
	stu.Person.name = "student a"
	stu.Person.age = 11
	stu.school = "yu xiu"

	fmt.Println("\n\n =============== stu ===============")
	fmt.Printf("%s, %d, %s", stu.name, stu.age, stu.school)

	// student1
	fmt.Println("\n\n =============== student1 ===============")
	student1 := Student{
		Person: Person{name: "wan hongbin ", age: 38},
		school: "fuzhou sixiao",
	}

	fmt.Println(student1)

	// 初始化2
	stu1 := Person{
		name: "张三",
		age:  10,
	}

	fmt.Println("\n\n =============== stu1 ===============")
	fmt.Printf("%s, %d", stu1.name, stu1.age)

	// student:不能提升字段
	var stu2 Student2

	stu2.person.name = "student a"
	//stu2.name = "student a"
	stu2.person.age = 11
	//stu2.age = 11
	stu2.school = "yu xiu"
}

// 1 定义父类
type Person struct {
	name string
	age  int
}

// 2  定义子类
type Student struct {
	Person
	school string
}

// 3  定义子类
type Student2 struct {
	person Person
	school string
}
