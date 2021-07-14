package main

import "fmt"

func main(){
	strs := []string{"cir", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	strLen := len(strs)
	if strLen == 0{
		return ""
	}
	if strLen == 1{
		return strs[0]
	}
	result := []byte(strs[0])
	len0 := len(result)
	if len0 == 0{
		return ""
	}
	maxIndex := len0 - 1
	for k := 1; k < len(strs); k++{
		vBytes := []byte(strs[k])
		len1 := len(vBytes)
		resultIndex := 0
		for i:= 0; i< len1; i++{
			if i == 0 && vBytes[i] != result[i]{
				return ""
			}
			if i < len0 && vBytes[i] == result[i]{
				if i <= maxIndex {
					resultIndex = i
				}
			}else{
				break
			}
		}
		maxIndex = resultIndex
	}
	re := maxIndex + 1
	return string(result[:re])
}
