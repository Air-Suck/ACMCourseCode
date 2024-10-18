//一道很明显的图论题，代码没什么好讲的，主要讲讲思路
//题目中的n就是结点的个数，k就是每个结点的最大度数，不妨假设有m条边
//首先m要大于等于n-1（因为树是边数最少的连通图）
//由握手定理，有度数和<=nk，所以2m<=nk，即m<=nk/2
//完全图的边数最多为n*（n-1）/2
//综上m满足不等式：n-1<=m<=min{nk/2,n*(n-1)/2}，其中n-1必<=n*(n-1)/2
package main

import (
	"fmt"
)

func main() {
	var (
		testNum int
		counNum int
		roadNum int
	)
	fmt.Scanf("%d\n", &testNum)
	for i := 0; i < testNum; i++ {
		fmt.Scanf("%d %d\n", &counNum, &roadNum)
        //如果n-1>nk/2时，m的取值范围为空集，即图不可能是连通图，输出0
		if 2*counNum-2 > counNum*roadNum {
			fmt.Printf("0\n")
			continue
		}
        //当min{nk/2,n*(n-1)/2}的值为n*(n-1)/2时，输出n*(n-1)/2
		if counNum*roadNum >= counNum*(counNum-1) {
			fmt.Printf("%d\n", counNum*(counNum-1)/2)
			continue
		} else {
            //当min{nk/2,n*(n-1)/2}的值为nk/2时，输出nk/2
			fmt.Printf("%d\n", counNum*roadNum/2)
		}
	}
}