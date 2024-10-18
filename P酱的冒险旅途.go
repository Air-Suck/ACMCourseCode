//就是一个统计字符，下面提供两份代码
package main

import (
	"fmt"
)

func main() {
	var (
		testNum int
		x       int
		y       int
		t       int
		charX   byte
		charY   byte
		i       int
	)
	fmt.Scanf("%d\n", &testNum)
	for j := 0; j < testNum; j++ {
		fmt.Scanf("%d %d %d\n", &x, &y, &t)
		input := make([]byte, t)
		for k, _ := range input {
			fmt.Scanf("%c", &input[k])
		}
		fmt.Scanf("\n")
        
		//特殊请款处理，下面一份代码一直wa就是这个原因
		if x == 0 && y == 0 {
			fmt.Printf("0\n")
			continue
		}
        //判断需要统计的字符
      	//x大于零则要向右，即统计R；小于零则要向左，即统计L
		if x >= 0 {
			charX = 'R'
		} else {
			charX = 'L'
            //将x置为正数用于计数
			x = 0 - x
		}
		//y方向情况同x方向
		if y >= 0 {
			charY = 'U'
		} else {
			charY = 'D'
			y = 0 - y
		}
        //遍历输入字符串
		for i = 0; i < t; i++ {
            //统计字符
			if input[i] == charX {
				x--
			}
			if input[i] == charY {
				y--
			}
            //如果已经能走出的时候就退出循环
			if y <= 0 && x <= 0 {
				break
			}
		}
        //上面的循环退出有两种情况，一种是完全遍历输入的字符串，这个时候i=t，这时还未到出口，应当输出-1；另一种情况是通过break退出，这个时候说明在遍历过程中已经到达出口，输出i+1
		if i == t {
			fmt.Printf("-1\n")
			continue
		} else {
			fmt.Printf("%d\n", i+1)
			continue
		}
	}
}

/*
	package main

	import (
		"fmt"
	)

	func main() {
		var (
			testNum int
			x       int
			y       int
			t       int
			target1 int
			target2 int
		)
		fmt.Scanf("%d\n", &testNum)
		for j := 0; j < testNum; j++ {
			fmt.Scanf("%d %d %d\n", &x, &y, &t)
			input := make([]byte, t)
			for i, _ := range input {
				fmt.Scanf("%c", &input[i])
			}
			fmt.Scanf("\n")
			//统计x方向的字符
			if x == 0 {
				target1 = 0
			} else if x > 0 {
				for target1 = 0; target1 < t; target1++ {
					if input[target1] == 'R' {
						x--
					}
					if x == 0 {
						break
					}
				}
				//当遍历一遍输入字符串仍然未找到足够的字符，说明无法走出，输出-1
				if target1 == t {
					fmt.Printf("-1\n")
					continue
				}
			} else {
				for target1 = 0; target1 < t; target1++ {
					if input[target1] == 'L' {
						x++
					}
					if x == 0 {
						break
					}
				}
				if target1 == t {
					//当遍历一遍输入字符串仍然未找到足够的字符，说明无法走出，输出-1
					fmt.Printf("-1\n")
					continue
				}
			}
			//统计y方向字符
			if y == 0 {
				target2 = 0
			} else if y > 0 {
				for target2 = 0; target2 < t; target2++ {
					if input[target2] == 'U' {
						y--
					}
					if y == 0 {
						break
					}
				}
				if target2 == t {
					//当遍历一遍输入字符串仍然未找到足够的字符，说明无法走出，输出-1
					fmt.Printf("-1\n")
					continue
				}
			} else {
				for target2 = 0; target2 < t; target2++ {
					if input[target2] == 'D' {
						y++
					}
					if y == 0 {
						break
					}
				}
				if target2 == t {
					//当遍历一遍输入字符串仍然未找到足够的字符，说明无法走出，输出-1
					fmt.Printf("-1\n")
					continue
				}
			}
			//要特判！！！！！
			if target1==0&&target2==0{
				fmt.Printf("0\n")
				continue
			}else{
				if target1 > target2 {
					fmt.Printf("%d\n", target1+1)
					continue
				} else {
					fmt.Printf("%d\n", target2+1)
					continue
				}   
			}
		}
	}
*/