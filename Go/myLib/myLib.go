package myLib

type X int
type Num struct {
	n, d int
}
func Mul(a, b Num) Num {
	a.n *= b.n; a.d *= b.d
	return a
}
