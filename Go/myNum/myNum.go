package myNum

type Num struct {
	N, D int	// numerator, denominator
}
func IniV(x *Num) {
	x.N = 0; x.D = 1
}

func gcd(a, b int) int {
	var c int
	for ; ; {
		if c = b % a; c == 0{ break }
		b = a; a = c
	}
	return a;
}
func PreProcess(n *Num) {
	if n.N == 0 {
		n.D = 1
	} else if n.D == 0 {
		n.N = 1
	}
	if(n.N & n.D) != 0 {
		x := gcd(n.N, n.D)
		n.N /= x; n.D /= x
	}
	if n.D < 0 {
		n.N *= -1; n.D *= -1
	}
}
func Calc(a, b Num, s byte) Num {
	var n Num
	if s == '+'{
		n.N = b.D * a.N + a.D * b.N
		n.D = a.D * b.D
	} else if s == '-' {
		n.N = b.D * a.N - a.D * b.N
		n.D = a.D * b.D
	} else if s == '*' {
		n.N = a.N * b.N
		n.D = a.D * b.D
	} else if s == '/' {
		n.N = a.N * b.D
		n.D = a.D * b.N
	}
	PreProcess(&n)
/*	if n.N == 0 { n.D = 1 }
	if (n.N & n.D) != 0{
		x = gcd(n.N, n.D)
		n.N /= x; n.D /= x
	}
	if n.D < 0 {
		n.N *= -1
		n.D *= -1
	}*/
	return n
}

func MinI(A []Num, n int) int {
	s := A[0]
	var p int = 0
	for i := 1; i < n; i++ {
		if s.N*A[i].D > s.D*A[i].N {
			s.N = A[i].N; s.D = A[i].D
			p = i
		}
	}
	return p
}
