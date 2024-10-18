#define _CRT_SECURE_NO_WARNINGS
#include <stdio.h>
#include<string.h>

char input[105][105] = { 0 };
char target[25] = { 0 };

int main() {
	int testNum, strNum, index = 0, spaceIndex = 0, i, j;

	scanf("%d", &testNum);
	for (i = 0;i < testNum;i++) {
		//record用于记录有用的论文数
		int record = 0;
		scanf("%d", &strNum);
		scanf("%s", target);
		while (getchar() != '\n');
		for (j = 0;j < strNum;j++) {
			fgets(input[j], 100, stdin);
			//fgets函数会读取末尾的换行符，所以将最后一位改为\0
			input[j][strlen(input[j]) - 1] = '\0';
		}
		for (j = 0;j < strNum;j++) {
			//flag用于记录当前论文名中是否含有关键词，因为一个论文只记录一次
			int flag = 0;
			//通过strcmp函数来进行比较，所以要将每一个空格改为\0来比较当前的单词是否是关键词
			index = 0;
			spaceIndex = 0;
			while (1) {
				//寻找下一个空格的位置（最后一个单词spaceIndex将到达\0）
				while (input[j][spaceIndex] != ' ' && input[j][spaceIndex] != '\0') {
					spaceIndex++;
				}
				//当已经找到最后一个词时，在循环外单独判断
				if (input[j][spaceIndex] == '\0') {
					break;
				}
				//将空格替换为结束标志，为调用strcmp函数做准备
				input[j][spaceIndex] = '\0';
				if (strcmp(target, &input[j][index]) == 0) {
					flag = 1;
					break;
				}
				spaceIndex++;
				index = spaceIndex;
			}
			if (strcmp(target, &input[j][index]) == 0) {
				flag = 1;
			}
			if (flag == 1) {
				record++;
			}
		}
		if (record != 0) {
			printf("%d\n\n", record);
		}
		//record等于0的时候要单独判断
		else {
			printf("Do not find\n\n");
		}
	}
	return 0;
}