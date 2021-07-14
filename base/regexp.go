package main 

import(
	"fmt"
	"regexp"
)

func main(){
	
	//查找字符串
	// find()

	//查找图片<img src='xxx'>并替换为超链接的图片<a href='xxx' ><img src='xxx'></a>
	findImg()

}

func findImg(){
	str := `<p>abcd</p><p>ABCD</p><p><img src="//s-static.yingxiong.com/file/1587455082758_99406450.png"></p><p>啊啊啊</p><p>哦哦哦</p>`

	reg := regexp.MustCompile("<img\\w>").FindAllString(str, -1)

	fmt.Println(str, reg)
}

func find(){
	str := "微信wx11123456我"
	reg := regexp.MustCompile("\\d").FindAllString(str, -1)
	fmt.Printf("%T\n", reg)//[]string
	fmt.Println(reg, len(reg))//[1 1 1 2 3 4 5 6] 8

	tmpMap := make(map[string]int)
	for _, v := range reg {
		_, isOk := tmpMap[v]
		if isOk != true {
			tmpMap[v] = 1
		}
	}

	fmt.Println(tmpMap, len(tmpMap))
}