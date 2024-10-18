package main

import "fmt"

func main() {
	var (
		testNum int
		data    string
		recR    = 0
		recY    = 0
		recB    = 0
		recL    = 0
	)
	//养成scanf里面塞\n的习惯
	fmt.Scanf("%d\n", &testNum)
	for testNum != 0 {
		//小心go语言的scanf会接收换行！！
		fmt.Scanf("%s\n", &data)
		recB = 0
		recL = 0
		recR = 0
		recY = 0
		//len函数可以计算字符串的长度
		for j := 0; j < len(data); j++ {
			if data[j] == 'R' {
				recR++
				//else一定要跟上一个大括号同行，否则会报语法错误
			} else if data[j] == 'Y' {
				recY++
			} else if data[j] == 'B' {
				recB++
			} else {
				recL++
			}
		}
		if recB != 0 {
			if recR == 7 {
				fmt.Println("Red")
			} else {
				fmt.Println("Yellow")
			}
		}
		if recL != 0 {
			if recY == 7 {
				fmt.Println("Yellow")
			} else {
				fmt.Println("Red")
			}
		}
		fmt.Scanf("%d\n", &testNum)
	}
}