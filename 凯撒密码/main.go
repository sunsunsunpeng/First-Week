package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scanf("%d\n",&n)
	var str string
	fmt.Scanf("%s",&str)
	for _,k:=range str{
		if k>='a'&&k<='z' {
			fmt.Printf("%s", string((k+int32(n)-'a')%26+'a'))
		}
	}
}