//难想.....
//思路如下：假设现在已经能表示的范围是 【1，P】，再加一个数Q（Q>P)，新增的可表示的范围为 【Q-P，P+Q 】，为了使范围尽可能的大，即有Q-P =P+1；即 Q=2P+1；新表示的范围即为【1，3P+1】。
package main

import (
	"fmt"
)

func main() {
	var (
		totNum  int
		i       int
		front   int
		current int
	)
	fmt.Scanf("%d\n", &totNum)
	record := make([]int, totNum+1)
	record[1] = 1
    //current表示当前元素应该需要的钱币种类
	current = 2
    //front表示下一次情况开始的第一个元素下表
	front = 2
    //死循环直到record[totNum]有值
	for {
        //从front到3*（front-1）+1都要赋值为current
		for i = front; i <= 3*front-2; i++ {
            //由于在循环外未进行totNum的判定，所以这里先进行判定
			if i > totNum {
				break
			}
			record[i] = current
		}
        //record[totNum]已经有值了
		if i > totNum {
			break
		}
        //每次循环过后钱币种类都要加一
		current++
        //front要变成3*（front-1）+1+1，紧接在上一次赋值的元素后继续赋值
		front = 3*front - 1
	}
	fmt.Printf("%d\n", record[totNum])
}