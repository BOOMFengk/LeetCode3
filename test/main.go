package main

import "fmt"






func main() {
	a:=" hello  world "
	 fmt.Println(reverseWords(a))

	



}
//i=1 slow =0 
func reverseWords(s string) string {
	ss:=[]byte(s)

	slow := 0 
	for i := 0; i < len(ss); i++ {
		if ss[i]!=' '{
			if slow!=0{
				ss[slow]=' '
				slow++
			}
			for i<len(s)&&ss[i]!= ' '{
				ss[slow]=ss[i]
				slow++
				i++
			}
		}
		
	}
	ss=ss[0:slow]


	j:=0
	reverseString(ss)
	for i := 0; i < len(ss); i++ {
		
		if ss[i]== ' '{
			reverseString(ss[j:i])
			j=i+1
		}else if i==len(ss)-1{
			reverseString(ss[j:i+1])
		}
	}
	return string(ss)



}
func reverseString(s []byte)  {
	left,right:=0,len(s)-1
	for left < right {
		s[left],s[right] = s[right],s[left]
		left++
		right--
	}

}

func reverseLeftWords(s string, n int) string {
        ss:=[]byte(s)

        reverseString(ss[0:n])
        reverseString(ss[n:len(ss)])
        reverseString(ss)
        return string(ss)
}

func reverseString(s []byte)  {
	left,right:=0,len(s)-1
	for left < right {
		s[left],s[right] = s[right],s[left]
		left++
		right--
	}
}