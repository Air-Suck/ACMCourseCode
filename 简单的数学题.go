//这真的是一道很简单的数学题
//首先因为负数也可以取模，所以可以去掉绝对值
//其次要证明，如果一个数是9的倍数，那么它的位和也为9的倍数，证法如下
//对正整数进行数学归纳（负整数同理）
//k=0时，N=9k=0，k=1时，N=9k=9，显然成立
//假设k=n时成立，k=n+1时，9N+9对9N的个位做枚举（进位的话会发现每次进位都会使位和减8），可知9N+9的位和也是9的倍数，所以只需要计算N-M的位和即可
//下面要证明N-M的位和就等于N的位和减去M的位和
//对N进行考虑，假设某位上N的数为k1，M的数为k2，那么要分两种情况（计算N的位和变化）
//若k1大于等于k2，那么N的位和的变化为-k2
//若k1小于k2，那么需要借位，则N的位和变化为：-1（借位）+k1+10-k2（借位后当前位的数）-k1=9（可忽略，因为是对9取模）-k2，即位和变化对9取模后为-M的位和
//所以N-M的位和即为N的位和减去M的位和，由于M是由N重组得到，所以位和差为0，即二者作差一定能被9整除
//那么问题就变成了求数字与N相同的最小整数M（一个快排搞定）
package main

import (
	"fmt"
)

// 升序排列（小的在前）
func quickSort(str []byte, front int, tail int) {
	left := front
	right := tail
	//选择基准值
	point := (front + tail) / 2
	flag := str[point]
	//当前组的元素个数小于等于0时退出当前循环
	if front >= tail {
		return
	}
	//当前组元素大于等于2时进行排序
	for front < tail {
		if str[front] < flag {
			front++
			continue
		}
		if str[tail] > flag {
			tail--
			continue
		}
		if str[tail] <= flag && str[front] >= flag {
			//交换两个数的值
			str[front], str[tail] = str[tail], str[front]
			front++
			tail--
			continue
		}
	}
	//优化，防止front或tail不动导致组规模不缩小造成死循环
	if front == left {
		front++
	}
	if tail == right {
		tail--
	}
	//递归
	quickSort(str, left, tail)
	quickSort(str, front, right)
}

func main() {
	i := 0
	var inputNum string
	input := make([]byte, 1001)
	fmt.Scanf("%s", &inputNum)
    //将读入的字符串存入切片，准备进行快排，退出循环时i就是最后一位的角标
	for i, _ = range inputNum {
		input[i] = inputNum[i]
	}
	//快排求M
	quickSort(input, 0, i)
    //输出M
	for j := 0; j <= i; j++ {
		fmt.Printf("%c", input[j])
	}
}