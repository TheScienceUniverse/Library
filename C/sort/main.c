#include <stdio.h>
#include <stdlib.h>
#include "display.h"
//#include "insertion.h"
//#include "selection.h"
#include "bubble.h"
int main(int argc, char *argv[]) {
	int i, n, *A;
	printf("Enter size of Array: ");
	scanf("%d", &n);
	A = (int*)malloc(n*sizeof(int));
	printf("Enter %d elements\n", n);
	for(i=0; i<n; scanf("%d", &A[i++]));
	printf("Array you have entered\n");
	showArray(A, n);
	sort(A, n);
	printf("Sorted Array:\n");
	showArray(A, n);
	return 0;
}
