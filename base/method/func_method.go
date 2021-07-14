// go没有OO的概念，所以这种功能是用来实现类的成员方法的

package main 
import "fmt"


func main(){
	var c Circle
	c.radius = 1
	fmt.Println(c.getArea())
}


/*定义结构体*/
type Circle struct{
	radius float64//成员变量
}

// func (variable_name variable_data_type) function_name(params) [return_type]{}
//成员方法，
//和普通方法的唯一区别就是普通方法在关键字func后面没有括号，紧接function_name()
//而这个方法需要在func后面加个括号，里面把该方法的拥有者作为参数传进来，
//形参就替代OO中的this，OO只中取成员是this.xxx，此处就为形参.xxx
func (c Circle) getArea() float64{
	return 3.14 * c.radius * c.radius
}