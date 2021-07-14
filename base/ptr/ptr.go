
// 指针是一个概念，指通过内存地址来访问变量
// 声明指针(变量)，和声明普通变量一致，只不过多一个*号
// var var_name *var_type 
// * 的作用是指定该变量作为一个指针，
// var_name 存放的内容是一个内存地址
// 获取该内存地址所表示的值通过 *var_name来获取

package main
import "fmt"

func main() {

	var a int = 10 //声明实际变量
	var ip *int //声明指针变量
	ip = &a //把实际变量的内存地址赋给指针变量

	//a变量内存地址是(&a)
	fmt.Printf("&a=%x\n", &a)

	fmt.Printf("a=%d\n", a)

	//变量的值是(值是a的内存地址)
	fmt.Printf("ip=%x\n", ip)

	fmt.Printf("*ip=%d\n", *ip)

	fmt.Printf("*&a=%d\n", *&a)

	emptyPtr()
}

func emptyPtr(){
	var ptr *int

	fmt.Printf("空指针ptr=%x\n", ptr)//输出：0

	// fmt.Printf("空指针*ptr=%x\n", *ptr)//会报错，因为地址是0，找不到值

	if ptr == nil {
		fmt.Printf("ptr == nil为true\n")
	}else{
		fmt.Printf("ptr == nil为false\n")
	}
}

