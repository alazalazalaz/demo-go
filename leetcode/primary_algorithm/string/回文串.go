package main

func main(){
	isPalindrome("race a car")
}

func isPalindrome(s string) bool {
	var newS []int32
	for _, s1 := range s{
		if s1 >= 'a' && s1 <= 'z' || s1 >= '0' && s1 <= '9'{
			newS = append(newS, s1)
		}
		if s1 >= 'A' && s1 <= 'Z'{
			newS = append(newS, s1+32)//转小写
		}
	}
	n := len(newS)
	for i := 0; i< n/2; i++{
		if newS[i] != newS[n-i-1]{
			return false
		}
	}
	return true
}

