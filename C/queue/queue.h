#ifndef _QUEUE_H
#define _QUEUE_H

#include <stdio.h>
#include <stdlib.h>

#define Q_SUCCESS 0
#define Q_EMPTY 1
#define Q_FULL 2

int *Q;
int posQ = -1;	//...Current Position
int maxQ = 0;	//...Maximum Size
int stat = 0;	//...Process Status

void makeQ(int n) {
	Q = (int*)malloc(n*sizeof(int));
	maxQ = n;
	printf("Queue is ready to use\nMax Size: %d integers\n", maxQ);
}
void showQ() {
	if(posQ == -1) {
		perror("Queue EMPTY\n");
	} else {
		for(int i=0; i<=posQ; printf("%d ", Q[i++]));
		printf("\n");
	}
}
void queueQ(int x){
	if(posQ == maxQ-1) {
		stat = Q_FULL;
	} else {
		Q[++posQ] = x;
		stat = Q_SUCCESS;
	}
}
void dequeueQ(int *x) {
	if(posQ < 0){
		*x = 0, stat = Q_EMPTY;
	} else {
		*x = Q[0];
		stat = Q_SUCCESS;
		--posQ;
		for(int i=0; i<=posQ; i++) { Q[i] = Q[i+1]; }
	}
}
#endif
