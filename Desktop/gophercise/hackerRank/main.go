package main

import (
	"fmt"
	"strings"
)

func main(){
	var input string
	fmt.Scanf("%s\n",&input)
	fmt.Println("Input is:",input)
    count:=1


	for _,char:=range input{

		s:=string(char)
		if strings.ToUpper(s)==s{
			count++
		}
		
		}
	


fmt.Println(count)
	}