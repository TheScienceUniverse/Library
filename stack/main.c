#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

int main(int argc, char *argv[]) {
	int i, x = 0, n=3;
	createS(n);
	for(i=0; i<n; i++) {
		pushS(i);
		showS();
	}
	for(i=0; i<n; i++) {
		popS(&x);
		showS();
	}
	deleteS();
	return 0;
}
