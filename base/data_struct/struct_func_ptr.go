package main 
import "fmt"


func main(){
	
	var u user
	log(u)//{ 0 false}

	u.name = "小张"
	u.age = 10
	u.sex = true
	log(u) //{小张 10 true}

	u.setName("小张setname")
	log(u) //{小张 10 true}

	//这里编译的时候会隐式转换为：
	//&u.setAge(20)
	u.setAge(20)
	log(u) //{小张 20 true}

	u2 := new(user)
	u2.name = "小王"
	u2.age = 30
	fmt.Println("u2=", u2)
	u2.setAge(40)
	u2.setName("小王 set name")
	fmt.Println("u2 updated =", u2)

}

type user struct{
	name string
	age int
	sex bool
}

//!!!!此处的u为user的一个副本，所以不会改变实参u!!!!
func (u user) setName(name string){
	u.name = name
}

//此处的u为user的一个引用，会改变实参u
func (u *user) setAge(age int){
	u.age = age
}


func log(u user){
	fmt.Printf("%v \n", u)
}

