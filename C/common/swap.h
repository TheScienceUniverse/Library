#ifndef _SWAP_H
#define _SWAP_H
void swap(int *a, int *b) {
	int t;
	t=*a, *a=*b, *b=t;
}
#endif
