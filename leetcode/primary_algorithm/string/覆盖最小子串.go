package main

import "fmt"

func main(){
	fmt.Println("结果:", minWindow("ADOBECODEBANC", "ABC"))
}

//滑动窗口
func minWindow(s string, t string) string {
	resultSlice := make([]byte, 0)
	sSlice := []byte(s)
	tSlice := []byte(t)
	sLen := len(sSlice)
	tMap,tMap2 := make(map[byte]int), make(map[byte]int)
	for i := range tSlice{
		tMap[tSlice[i]]++
		tMap2[tSlice[i]]++
	}
	//先找到第一个窗口
	left := 0
	findFirst := false
	for right := 0; right<sLen; right++{
		if findFirst == false{
			if _, isExist := tMap[sSlice[right]]; isExist == true{
				if tMap[sSlice[right]] == 1{
					delete(tMap, sSlice[right])
				}else{
					tMap[sSlice[right]]--
				}
			}

			if len(tMap) == 0{
				findFirst = true
				resultSlice = sSlice[left:right+1]
			}
		}else{
			if _, isExist := tMap2[sSlice[right]]; isExist == true{
				resultSlice = sSlice[left:right+1]
				for _check(resultSlice, t) == true {
					left++
					resultSlice = sSlice[left:right+1]
				}
				left--
				resultSlice = sSlice[left:right+1]
			}
		}
	}
	return string(resultSlice)
}

func _check(resultSlice []byte, t string) bool {
	tSlice:= []byte(t)
	tMap2:= make(map[byte]int)
	for i := range tSlice{
		tMap2[tSlice[i]]++
	}
	for i:=0; i<len(resultSlice); i++{
		if _, isExist := tMap2[resultSlice[i]]; isExist == true{
			if tMap2[resultSlice[i]] == 1{
				delete(tMap2, resultSlice[i])
			}else{
				tMap2[resultSlice[i]]--
			}
		}
	}
	if len(tMap2) == 0{
		return true
	}
	return false
}
