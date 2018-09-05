package myGraph
import(
	"fmt"
mN	"../myNum"
)
type Point struct {
	X, Y int
}
type LS struct {	// Line Segment
	P1, P2 Point
}
func IPoint(l1, l2 LS) (Point, string) {	// intersection point
	var msg string = "OK"
	p := Point{0, 0}
	var m1, m2 mN.Num
	m1.N = (l1.P2.Y - l1.P1.Y); m1.D = (l1.P2.X - l1.P1.X)
	m2.N = (l2.P2.Y - l2.P1.Y); m2.D = (l2.P2.X - l2.P1.X)
	mN.PreProcess(&m1); mN.PreProcess(&m2)
	if (m1.N == m2.N) && (m1.D == m2.D) {
		msg = "ERROR!...Lines Parallel/Overlapping"
	} else {

	}
	return p, msg
}
