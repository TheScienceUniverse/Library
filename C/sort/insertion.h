#ifndef _INSERTION_H
#define _INSERTION_H
#include "myAlgo.h"
void sort(int *A, int n) {
	int i, j;
	for(i=0; i<n; i++)
		printf("%d ", A[i]);
	printf("\n");

	i=1;
	while(i<n){
		j=i;
		while (j>0 && A[j-1]>A[j]) {
			swap(&A[j], &A[j-1]);
			--j;
		}
		++i;
	}
}
#endif
