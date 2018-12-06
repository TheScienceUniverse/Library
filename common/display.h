#ifndef _DISPLAY_H
#define _DISPLAY_H
void dsplArr1Dint (int *A, int n) {
	for(int i=0; i<n; printf("%d ", A[i++]));
	printf("\n");
}
void dsplArr2Dint (int *A, int r, int c) {
	for(int i=0; i<r; i++) {
		for (int j=0; j<c; printf("%d ", A[i][j++]));
		printf("\n");
	}
}
#endif
