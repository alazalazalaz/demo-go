package main

import (
	"fmt"
	"math"
)

//洋葱模型
//一层一层进入洋葱内部，一层一层的就是middleware，每次调用middleware的next()就表示进入下一层，洋葱内部就是get/post等路由
//到达内部之后，再一层一层的返回第二遍执行middleware next()后面的内容
const maxIndex = math.MaxInt8 / 2

type myFuncHandle func(ctx *myContext)
type myContext struct{
	routeArray []myFuncHandle
	index int8
}

func (m *myContext) next(){
	if m.index < maxIndex{
		m.index++
		m.routeArray[m.index](m)
	}
}

func (m *myContext) abort(){
	m.index = maxIndex
	fmt.Println("已被终止...")
}

func (m *myContext) use(f myFuncHandle){
	m.routeArray = append(m.routeArray, f)
}

func (m *myContext) get(uri string, f myFuncHandle){
	m.routeArray = append(m.routeArray, f)
}

func (m *myContext) run(){
	m.routeArray[0](m)
}

func main(){
	ctx := &myContext{}
	ctx.use(m1)
	ctx.use(m2)
	ctx.get("hahahah", getF)
	ctx.run()
}

func m1(ctx *myContext){
	fmt.Println("im m1 begin")
	ctx.abort()
	ctx.next()
	fmt.Println("im m1 end")
}

func m2(ctx *myContext){
	fmt.Println("im m2 begin")
	ctx.next()
	fmt.Println("im m2 end")
}

func getF(ctx *myContext){
	fmt.Println("im get function")
}

