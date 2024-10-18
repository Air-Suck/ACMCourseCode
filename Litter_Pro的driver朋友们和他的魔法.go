// 本题审题出错哩，下面是只能交换相邻driver的版本
package main

import (
	"fmt"
)

func main() {
	var (
		driverNum int
		record1   int
		record2   int
	)
	fmt.Scanf("%d\n", &driverNum)
	queue1 := make([]byte, driverNum)
	queue2 := make([]byte, driverNum)
    //初始化两个切片来记录0开头的情况和1开头的情况
	for i := 0; i < driverNum; i++ {
		fmt.Scanf("%c", &queue1[i])
		queue2[i] = queue1[i]
	}
    //对queue1进行变换
	for i := 1; i < driverNum-1; i++ {
        //当前元素和前一个元素不相同时，说明当前元素正确，进行下一次循环
		if queue1[i-1] != queue1[i] {
			continue
		} else {
            //当当前元素和前一个元素相同时，说明当前元素需要魔法，但是用交换还是变换取决于后一个元素
            //当前元素和下一个元素不相同时，说明当前元素和下一个元素都是错误的，所以用交换魔法
			if queue1[i+1] != queue1[i] {
				record1++
				queue1[i+1], queue1[i] = queue1[i], queue1[i+1]
                //注意到这里是一次将两个元素的值调整正确
                //所以i要加2（这里只加一次是因为退出当前次时i还会自增）
				i++
			} else {
                //当前元素和下一个元素相同时，说明当前元素错误，但是下一个元素正确，所以使用变换魔法
				record1++
				if queue1[i+1] == '1' {
					queue1[i]--
				} else {
					queue1[i]++
				}
                //同样的，这里也是一次改正两个元素的值，所以i要加二
				i++
			}
		}
	}
    //由于循环中使用到了i+1，所以最后一个元素没办法在循环中处理，所以在循环外处理最后一个元素
	if queue1[driverNum-1] == queue1[driverNum-2] {
		record1++
		if queue1[driverNum-2] == '1' {
			queue1[driverNum-1] = '0'
		} else {
			queue1[driverNum-1] = '1'
		}
	}
    //对queue2进行首元素的调整
    //如果第一个元素和第二个元素不相同的话，说明两个元素都是错误的，使用交换魔法
	if queue2[0] != queue2[1] {
		record2++
		queue2[0], queue2[1] = queue2[1], queue2[0]
	} else {
        //第一个元素和第二个元素相同时，说明第一个元素错误，第二个元素正确，所以使用变换魔法
		record2++
		if queue2[0] == '1' {
			queue2[0] = '0'
		} else {
			queue2[0] = '1'
		}
	}
    //通过对queue2的预处理使其前两个元素都是正确的，所以i从2开始，下面的操作流程与queue1的操作完全相同
	for i := 2; i < driverNum-1; i++ {
		if queue2[i-1] != queue2[i] {
			continue
		} else {
			if queue2[i+1] != queue2[i] {
				record2++
				queue2[i+1], queue2[i] = queue2[i], queue2[i+1]
				i++
			} else {
				record2++
				if queue2[i+1] == '1' {
					queue2[i]--
				} else {
					queue2[i]++
				}
				i++
			}
		}
	}
	if queue2[driverNum-1] == queue2[driverNum-2] {
		record2++
		if queue2[driverNum-2] == '1' {
			queue2[driverNum-1] = '0'
		} else {
			queue2[driverNum-1] = '1'
		}
	}
    //输出queue1和queue2两种情况中的较小者
	if record1 <= record2 {
		fmt.Printf("%d", record1)
	} else {
		fmt.Printf("%d", record2)
	}
}

//看不懂代码所以就简单讲讲实现逻辑
//本质上还是贪心算法
//假设有两个队列，queue1存储0开头的情况（奇数位都是0），queue2存储1开头的情况（奇数位都是1）
//以queue1为例，计算奇数位上1的个数（计为oneNum）和偶数位上0的个数（计为zeroNum），这些都是不正确的数字
//假设oneNum<zeroNum，那么需要oneNum次交换魔法将奇数位上的1变为0
//此时还剩下zeroNum-oneNum个0在偶数位上（因为已经将oneNum个0和奇数位上的1交换了）
//这些数就需要执行变换魔法，所以总次数为oneNum（交换）+zeroNum-oneNum（变换）=zeroNum次
//最后输出queue1和queue2情况下的最小值
/*	#include <bits/stdc++.h>

	using namespace std;

	int magic(int a,int b)
	{
		return a>b? a:b;
	}

	int magic_sum(int a,int b)
	{
		return a<b? a:b;
	}

	int main()
	{
		int n,i,sum;
		int flag[4]={0};
		scanf("%d",&n);
		getchar();
		char c[n+1]={0};
		cin.getline(c,n+1);
		for(i=0;i<n;i++)
		{
			if(i%2)
			{
				if(c[i]=='1') flag[1]++;
				else flag[3]++;
			}
			else
			{
				if(c[i]=='0') flag[0]++;
				else flag[2]++;
			}
		}
		sum=magic_sum(magic(flag[0],flag[1]),magic(flag[2],flag[3]));
		cout << sum << endl;

		return 0;
	}
*/