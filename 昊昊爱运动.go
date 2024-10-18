//go语言代码过不了，但是算时间复杂度应该是不会超时的。本题最重要的应该是预处理
package main

import (
	"fmt"
)

func main() {
	var (
		dayNum     int
		sportsType int
		quesNum    int
		left       int
		right      int
		x          int
	)
	fmt.Scanf("%d %d\n", &dayNum, &sportsType)
	record := make([]int, dayNum+1)
    //创建一个二维切片用于预处理
	pro := make([][]int, dayNum+1)
	for i := 1; i <= dayNum; i++ {
		pro[i] = make([]int, dayNum+1)
	}
    
	for i := 1; i <= dayNum; i++ {
		fmt.Scanf("%d", &record[i])
	}
	fmt.Scanf("\n")
	//预处理，直接计算出从第i天到第j天运动的种类并存储到pro[i][j]中
	for i := 1; i <= dayNum; i++ {
		x = 0
        //通过一个数组来进行去重，便于统计运动种类
		hash := make([]int, sportsType+1)
		for j := i; j <= dayNum; j++ {
			if hash[record[j]] == 0 {
				x++
				hash[record[j]] = 1
			}
			pro[i][j] = x
		}
	}

	fmt.Scanf("%d\n", &quesNum)
	for i := 0; i < quesNum; i++ {
		fmt.Scanf("%d %d\n", &left, &right)
		fmt.Printf("%d\n", pro[left][right])
	}
}
/*
这里提供一份能过的cpp代码
	#include<cstdio>
	#include<iostream>
	#include<algorithm>
	#include<cstring>
	using namespace std;
	int n,m,i,tmp,j,Q,l,r,ans,a[2005][105];
	int main(){
		scanf("%d%d",&n,&m);
		for(i=1;i<=n;++i){
			scanf("%d",&tmp);
			a[i][tmp]=1;
			for(j=1;j<=m;++j)a[i][j]+=a[i-1][j];
		}
		scanf("%d",&Q);
		for(i=1;i<=Q;++i){
			scanf("%d%d",&l,&r);
			ans=0;
			for(j=1;j<=m;++j)ans+=(a[r][j]>a[l-1][j]);
			printf("%d\n",ans);
		}
		return 0;
	}
*/