package main
import (
	"fmt"
	"os"
C	"./des"
)
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("ERROR: No Input File")
		os.Exit(0)
	}
	fpi, err := os.Open(os.Args[1])
	check(err)
	fpo, err := os.Create("op.data")
        check(err)

	X := make([]byte, 8)

// Testing
n, err := fpi.Read(X)
check(err)

var i int
if(n != 0) {
	for i = n-1; i < 8; i++ {
		X[i] = 0x00;
	}
	fmt.Println(X)
	X = C.DES(X)
//	fmt.Println(X)
}

/*
	for ; ; {
		n, err := fpi.Read(X)
		check(err)
		if(n == 0) {
			break
		}

//		fmt.Println("#bytes: ", n)
//		fmt.Println(X)

		X = C.DES(X)

		n, err = fpo.Write(X)
		check(err)
//		fmt.Println("#bytes", n)
	}
*/
// Rewind
//	_, err = fpi.Seek(0, 0)
//	check(err)



	fpo.Close()
	fpi.Close()
}
