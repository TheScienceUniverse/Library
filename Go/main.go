package main

import (
	"fmt"
//mN	"./myNum"
mG	"./myGraph"
)

func main() {
	fmt.Println("Hello, World!\nWelcome to Sweep Line Algorithm to find intersecting points of line segments")
	var i, nL int
	fmt.Print("Enter the number of lines: ")
	fmt.Scanf("%d", &nL)
	G := make([]mG.LS, nL)
	fmt.Println("Array created to store end-points")
	var x1, y1, x2, y2 int = 0, 0, 0, 0
	for i = 0; i < nL; i++ {
		fmt.Println("Line No.",i+1)
		fmt.Println("Enter x1, y1, x2, y2:")
		fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2);
		G[i].P1.X, G[i].P1.Y, G[i].P2.X, G[i].P2.Y = x1, y1, x2, y2
	}
	fmt.Println(G)
	fmt.Println(mG.IPoint(G[0], G[1]))
	fmt.Println("Thank You for using it")
}
