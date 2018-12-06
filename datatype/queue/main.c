#include <stdio.h>
#include <stdlib.h>
#include "queue.h"

int main(int argc, char *argv[]) {
	int i, x = 0, n=3;
	createQ(n);
	for(i=0; i<n; i++) {
		enqueueQ(i);
		showQ();
	}
	for(i=0; i<n; i++) {
		dequeueQ(&x);
		showQ();
	}
	deleteQ();
	return 0;
}
