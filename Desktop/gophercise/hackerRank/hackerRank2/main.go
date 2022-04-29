package main

import (
	"fmt"
	"strings"
)

func main(){
var input string
var Len,word int
fmt.Scanf("%d\n",&Len)
fmt.Scanf("%s\n",&input)
fmt.Scanf("%d\n",&word)
str:=""
s:="abcdefghijklmnopqrstuvwxyz"
t:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"



for _,char:=range input{

	switch{
	case strings.IndexRune(s,char) >=0:
		str=str + string(convert(char,word,[]rune(s)))
	
    case strings.IndexRune(t,char) >=0:
	str=str + string(convert(char,word,[]rune(t)))
	
    default:
		str=str + string(char)

	}
}

	
	fmt.Println(str)
}





func convert(r rune,word int,key []rune)rune{
	index:=strings.IndexRune(string(key),r)


	if index<0{
		panic("index is less than zero")
	}
	index = (index+word) % len(key)
	return key[index]
}

