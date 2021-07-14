package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(isAnagram("abcd", "bdca"))
}

//什么是异位词
//长度一样，包含的字母一样，每个字符出现的频率也一样，只是顺序不同而已
//官方方法
//1、将s和l排序
//2、如果是异位词，排序后应该是完全相同的两个字符串！！！！
func isAnagram(s string, t string) bool {
	sliceS := []byte(s)
	sliceT := []byte(t)
	sort.Slice(sliceS, func(i, j int) bool {return sliceS[i] > sliceS[j]})
	sort.Slice(sliceT, func(i, j int) bool {return sliceT[i] > sliceT[j]})
	return string(sliceS) == string(sliceT)
}

//方法2.1
//方法2还可以优化，用一个字典够了，
//1、字典或者新数组存放s的频次
//2、for判断l，减去字典或新数组中的频次，
//3、如果是异位词，那字典或者新数组的value应该都为0，如果有>0或者<0的都不是异位词。
//方法2
//哈希，存放每个字母的频次，但是需要判断两次map
// func isAnagram(s string, t string) bool {
//     dictS, dictT := make(map[rune]int), make(map[rune]int)
//     for _, numS := range s{
//         dictS[numS]++
//     }
//     for _, numT := range t{
//         dictT[numT]++
//     }
//     for k, v :=range dictS{
//         if vT, isExist := dictT[k]; isExist == true && vT == v{

//         }else{
//             return false
//         }
//     }
//     for k, v :=range dictT{
//         if vS, isExist := dictS[k]; isExist == true && vS == v{

//         }else{
//             return false
//         }
//     }
//     return true
// }
