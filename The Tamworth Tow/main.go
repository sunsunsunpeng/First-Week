package main

import "fmt"

var (
	Map [11][11]byte
    fx,fy,cx,cy int       //农夫，牛所处位置
	Fx,Fy,Cx,Cy int       //农夫，牛移动后的方向,x为行，y为列
	f=1;c=1               //农夫，牛移动的方向,1,2,3,4,北，东，南，西
	sum=0                 //计时总和
)
func main(){
	for i:=1;i<=10;i++{
		for j:=1;j<=10;j++{
			if j==10 {
				fmt.Scanf("%c\n", &Map[i][j])
			}else{
				fmt.Scanf("%c", &Map[i][j])
			}
			if Map[i][j] == 'F'{
				fx=i
				fy=j
				Map[i][j]='.'
			}
			if Map[i][j] == 'C'{
				cx = i
				cy = j
				Map[i][j]='.'
			}
		}
	}
		for {
			MoveF()
			MoveC()
			sum++
			if (fx==cx)&&(fy==cy){
				fmt.Println(sum)
				break
			}
			if sum>10000{
				fmt.Println(0)
				break
			}
		}
}
func MoveF(){
	if f==1{
		Fx=fx-1
		Fy=fy
	}else if f==2{
		Fx=fx
		Fy=fy+1
	}else if f==3{
		Fx=fx+1
		Fy=fy
	}else if f==4{
        Fx=fx
		Fy=fy-1
	}
	if (Fx >= 1) && (Fy >= 1) && (Fy <= 10) && (Fx <= 10) && (Map[Fx][Fy] == '.'){
		fx=Fx
		fy=Fy
	}else{
		f++
		if f>4{
			f=1
		}
	}
}
func MoveC(){
	if c==1{
		Cx=cx-1
		Cy=cy
	}else if c==2{
		Cx=cx
		Cy=cy+1
	}else if c==3{
		Cx=cx+1
		Cy=cy
	}else if c==4{
		Cx=cx
		Cy=cy-1
	}
	if (Cx >= 1) && (Cy >= 1) && (Cy <= 10) && (Cx <= 10) && (Map[Cx][Cy] == '.'){
		cx=Cx
		cy=Cy
	}else{
		c++
		if c>4{
			c=1
		}
	}
}