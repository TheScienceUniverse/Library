#include <stdio.h>
#include <stdlib.h>
#include "queue.h"

int main(int argc, char *argv[]) {
	int i, x = 0, n=3;
	makeQ(n);
	for(i=0; i<n; i++) {
		queueQ(i);
		showQ();
	}
	for(i=0; i<n; i++) {
		dequeueQ(&x);
		showQ();
	}

	return 0;
}
