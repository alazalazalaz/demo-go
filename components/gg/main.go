package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"math"
)

const WIDTH = 600
const HEIGHT = 400
const LINE_SPACE = 5

//测试gg库，gg库是用golang来绘制2D图形的库，还能添加文字。
func main(){

	//生成一张图片
	genImage()

	//生成一张图片并写字
	genImageWrite()

	//生成一张图片并写字并且让字自动换行，到末尾时自动补充省略号
	genImageWriteV2()
}

func genImage(){
	dc := gg.NewContext(100, 100)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	if err := dc.SavePNG("./components/gg/data/genImage.png"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("genImage finish!")
}

func genImageWrite(){
	const WIDTH = 600
	const HEIGHT = 400
	var fontPath = "./components/gg/font/JiZiJingDianJiSongJianFan.ttf"
	var savePath = "./components/gg/data/genImageWrite.png"
	var textString = "hello worldabcdefghijklmn"
	dc := gg.NewContext(WIDTH, HEIGHT)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(fontPath, 100); err != nil {
		log.Fatalln(err)
	}
	dc.DrawString(textString, 10, 100)

	if err := dc.SavePNG(savePath); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(savePath + " finish!")
}

func genImageWriteV2(){
	var fontPath = "./components/gg/font/JiZiJingDianJiSongJianFan.ttf"
	var savePath = "./components/gg/data/genImageWriteV2.png"
	var textString = "一二三\n四五六七八九十我说这是最后一行了abcdef"
	dc := gg.NewContext(WIDTH, HEIGHT)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(fontPath, 10); err != nil {
		log.Fatalln(err)
	}

	var currentX, currentY float64 = 0, 10
	_handleOneText(dc, textString, &currentX, &currentY)

	if err := dc.SavePNG(savePath); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("string len=%d, byte=%d, rune=%d\r\n", len(textString), len([]byte(textString)), len([]rune(textString)))
	fmt.Println(savePath + " finish!")
}

func _handleOneText(dc *gg.Context, text string, currentX *float64, currentY *float64){
	//计算当前字符串长度和宽度
	sWidth, sHeight := dc.MeasureString(text)
	fmt.Printf("fontHeight:%0.2f, current text size:%0.2f, %0.2f \r\n", dc.FontHeight(), sWidth, sHeight)
	//当前字符串需要几行才能显示完
	currentTextMaxLineNum := int(math.Ceil(sWidth/WIDTH))
	fmt.Printf("currentTextMaxLineNum:%d\n", currentTextMaxLineNum)

	//计算图片最多能容纳几行
	imageMaxLineNum := int(HEIGHT/(sHeight + LINE_SPACE))
	fmt.Printf("imageMaxLineNum:%d\n", imageMaxLineNum)

	//将字符串分割成每行数组
	textRune := []rune(text)
	tmpStrRune := make([]rune, 0)
	textArray := make([]string, 0)
	for _, oneWord := range textRune{
		//\r\n识别为换行符
		if oneWord == 10 || oneWord == 13{
			//主动换行
			textArray = append(textArray, string(tmpStrRune))
			tmpStrRune = nil
			continue
		}
		tmpStrRune = append(tmpStrRune, oneWord)
		if wid, _ := dc.MeasureString(string(tmpStrRune)); wid > WIDTH {
			//被动换行
			//删除最后一个元素
			lastOneWord := tmpStrRune[len(tmpStrRune) - 1]
			tmpStrRune = tmpStrRune[:len(tmpStrRune) - 1]
			//换行
			textArray = append(textArray, string(tmpStrRune))
			tmpStrRune = nil
			tmpStrRune = append(tmpStrRune, lastOneWord)
		}
	}

	if len(tmpStrRune) > 0 {
		textArray = append(textArray, string(tmpStrRune))
	}

	//截取要显示的行数，并且如果超出，最后一行最后三个字替换为省略号...
	if len(textArray) > imageMaxLineNum {
		textArray = textArray[:imageMaxLineNum]
		displayLastLine := []rune(textArray[imageMaxLineNum - 1])
		textArray[imageMaxLineNum - 1] = string(displayLastLine[:len(displayLastLine) - 3]) + "..."
	}

	//打印字符串数组
	for _, v := range textArray{
		dc.DrawString(v, *currentX, *currentY)
		*currentY += dc.FontHeight() + LINE_SPACE
	}

}
