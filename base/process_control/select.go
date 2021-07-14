package main

import "fmt"

func main(){
	/**
	select和switch的语法一致，只不过case后面只能接channel的收发操作
	 */

}

/**
下面这个控制结构会等待 c <- x 或者 <-quit 两个表达式中任意一个返回。
无论哪一个表达式返回都会立刻执行 case 中的代码，当 select 中的两个 case 同时被触发时，会随机执行其中的一个。

当我们在 Go 语言中使用 select 控制结构时，会遇到两个有趣的现象：
select 能在 Channel 上进行非阻塞的收发操作；
select 在遇到多个 Channel 同时响应时，会随机执行一种情况；
 */
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}