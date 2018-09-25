#ifndef _SELECTION_H
#define _SELECTION_H
#include "myAlgo.h"
int *sort(int *A, int n) {
	int i, j, iMin;
	for(i=0; i<n-1; i++) {
		iMin = i;
		for(j=i+1; j<n; j++) {
			if(A[j]<A[iMin])
				iMin=j;
		}
		if(iMin!=i)
			swap(&A[i], &A[iMin]);
	}
	return A;
}
#endif
