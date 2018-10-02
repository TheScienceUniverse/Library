#ifndef _STACK_H
#define _STACK_H

#include <stdio.h>
#include <stdlib.h>

#define S_SUCCESS 0
#define S_EMPTY 1
#define S_FULL 2

int *S;
int posS = -1;	//...Current Position
int maxS = 0;	//...Maximum Size
int stat = 0;	//...Process Status

void createS(int n) {
	S = (int*)malloc(n*sizeof(int));
	maxS = n;
	printf("Queue is ready to use\nMax Size: %d integers\n", maxS);
}
void showS() {
	if(posS == -1) {
		perror("Queue EMPTY\n");
	} else {
		for(int i=0; i<=posS; printf("%d ", S[i++]));
		printf("\n");
	}
}
void deleteS() {
	free(S);
	posS = -1, maxS = 0;
	stat = S_SUCCESS;
}
void pushS(int x){
	if(posS == maxS-1) {
		stat = S_FULL;
	} else {
		S[++posS] = x;
		stat = S_SUCCESS;
	}
}
void popS(int *x) {
	if(posS < 0){
		*x = 0, stat = S_EMPTY;
	} else {
		*x = S[posS--];
		stat = S_SUCCESS;
	}
}

#endif
