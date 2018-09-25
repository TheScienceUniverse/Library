#ifndef _BUBBLE_H
#define _BUBBLE_H
#include "myAlgo.h"
void sort(int *A, int n) {
	int i, j;
	for (i = 0; i < n-1; i++) {
		for (j = 0; j < n-i-1; j++)
			if (A[j] > A[j+1])
				swap(&A[j], &A[j+1]);
	}
}
#endif
