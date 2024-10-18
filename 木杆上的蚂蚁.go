//总体思路一样还是蚂蚁直接穿过不掉头。每次只要找到向左或者向右时间最短的蚂蚁即可
//如果时间最短的蚂蚁是向左走的，那就输出最左边的一只蚂蚁
//如果时间最短的蚂蚁是向右走的，那就输出最右边的一只蚂蚁
package main

import "fmt"

type antData struct {
	name      string
	direction byte
	time      int
}

//快排用于对time进行排序
// 升序排列（小的在前）
func quickSort(str []antData, front int, tail int) {
	left := front
	right := tail
	//选择基准值
	point := (front + tail) / 2
	flag := str[point].time
	//当前组的元素个数小于等于0时退出当前循环
	if front >= tail {
		return
	}
	//当前组元素大于等于2时进行排序
	for front < tail {
		if str[front].time < flag {
			front++
			continue
		}
		if str[tail].time > flag {
			tail--
			continue
		}
		if str[tail].time <= flag && str[front].time >= flag {
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
	var (
		testNum  int
		stickLen int
		antNum   int
		position int
		index    int
		left     int
		right    int
	)

	fmt.Scanf("%d\n", &testNum)
	for i := 1; i <= testNum; i++ {
		fmt.Scanf("%d %d\n", &antNum, &stickLen)
        //stick切片模拟木杆
		stick := make([]antData, stickLen+1)
        //queue切片通过哈希将蚂蚁从左到右排序
		queue := make([]antData, antNum+1)
        //time来记录最短时间
		time := make([]antData, antNum+1)
		//木杆的初始化
		for j := 0; j < antNum; j++ {
			T := new(antData)
			fmt.Scanf("%s %d %c\n", &(T.name), &position, &(T.direction))
			if T.direction == 'L' {
				T.time = position
			} else {
				T.time = stickLen - position
			}
			stick[position] = *T
		}
		//预处理（给queue和time赋值）
		index = 0
		for j := 0; j < stickLen; j++ {
            //跳过空位置
			if stick[j].direction == 0 {
				continue
			} else {
				queue[index] = stick[j]
				time[index] = stick[j]
				index++
			}
		}
		//给time排序
		quickSort(time, 0, antNum-1)
		//开始遍历time和queue
		left = 0
		right = antNum - 1
		fmt.Printf("Case #%d:\n", i)
        //总共要掉下来antNum只蚂蚁，所以循环次数为antNum
		for j := 0; j < antNum; j++ {
            //如果时间最短的蚂蚁是向左走的
			if time[j].direction == 'L' {
                //输出queue中最左端的蚂蚁
				fmt.Printf("%d %s\n", time[j].time, queue[left].name)
                //最左端的蚂蚁掉下木杆
				left++
			} else {
                //输出queue中最右端的蚂蚁
				fmt.Printf("%d %s\n", time[j].time, queue[right].name)
                //最右端的蚂蚁掉下木杆
				right--
			}
		}
	}
}