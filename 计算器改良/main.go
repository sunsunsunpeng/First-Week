package main

import (
    "fmt"
)

func main(){
    var (
    	str string
        unknown string
    )
    total:=0
    result:=0
    coefficient:=0
    Equ:=1
    tag:=1
    fmt.Scanf("%s",&str)
    for _,k:=range str{
    if k=='+'{
        result -= total * tag * Equ
        total = 0
        tag = 1
    }else if k == '-'{
        result -= total * tag * Equ
        total = 0
        tag = -1
    }else if k == '='{
        result -= total * tag * Equ
         Equ = -1
         total = 0
         tag = 1
    }else if ((k>='a')&&(k<='z'))||((k>='A')&&(k<='Z')){
        coefficient = coefficient + total * Equ*tag
        unknown = string(k)
        if coefficient == 0 {
            coefficient = 1
        };
        total = 0
    }else{
        total = total * 10 + int(k-'0')
    }
    }
    result += total * tag
    fmt.Printf("%s=%.3f",unknown,float64(result)/float64(coefficient))
}
