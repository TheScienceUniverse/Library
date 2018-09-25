#ifndef _DISPLAY_H
#define _DISPLAY_H
void showArray(int *A, int n) {
	for(int i=0; i<n; printf("%d ", A[i++]));
	printf("\n");
}
#endif
